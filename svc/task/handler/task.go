package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/haileemiu/manage-life/ent"
	"github.com/haileemiu/manage-life/ent/task"
	"github.com/haileemiu/manage-life/pkg/res"
	"github.com/haileemiu/manage-life/svc/task/model"
)

type Task struct {
	ent *ent.Client
}

func New(entClient *ent.Client) *Task {
	return &Task{ent: entClient}
}

func (t Task) Routes(r chi.Router) {
	r.Get("/", t.list)
	r.Post("/", t.create)
	r.Get("/{id}", t.getByID)
	r.Put("/{id}", t.update)
	r.Delete("/{id}", t.delete)
}

func (t Task) list(w http.ResponseWriter, r *http.Request) {
	tasks, err := t.ent.Task.Query().All(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var taskList []model.TaskItemResponse
	for _, entModel := range tasks {
		task := model.TaskItemResponse{
			ID: 						entModel.ID,
			Title:          entModel.Title,
			Notes:          entModel.Notes,
			IsTimeSenstive: entModel.IsTimeSenstive,
			IsImportant:    entModel.IsImportant,
			RemindAt:       entModel.RemindAt,
			DueAt:          entModel.DueAt,
		}
		taskList = append(taskList, task)
	}

	if err := json.NewEncoder(w).Encode(taskList); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
}

func (t Task) create(w http.ResponseWriter, r *http.Request) {
	req := model.TaskCreateRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if ok, errs := req.Validate(); !ok {
		res.NewValidationErrorResponse(errs).Send(w)
		return
	}

	entTask, err := t.ent.Task.Create().
		SetTitle(req.Title).
		SetNotes(req.Notes).
		SetIsTimeSenstive(req.IsTimeSenstive).
		SetIsImportant(req.IsImportant).
		Save(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	task := model.TaskItemResponse{
		ID:             entTask.ID,
		Title:          entTask.Title,
		Notes:          entTask.Notes,
		IsTimeSenstive: entTask.IsTimeSenstive,
		IsImportant:    entTask.IsImportant,
		RemindAt:       entTask.RemindAt,
		DueAt:          entTask.DueAt,
	}

	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
}

func (t Task) getByID(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	task, err := t.ent.Task.Query().Where(task.ID(taskID)).Only(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	taskItem := model.TaskItemResponse{
		ID:             task.ID,
		Title:          task.Title,
		Notes:          task.Notes,
		IsTimeSenstive: task.IsTimeSenstive,
		IsImportant:    task.IsImportant,
		RemindAt:       task.RemindAt,
		DueAt:          task.DueAt,
	}

	if err := json.NewEncoder(w).Encode(taskItem); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
}

func (t Task) update(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	req := model.TaskCreateRequest{}

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// TODO: a different validate for PUT? 
	// if ok, errs := req.Validate(); !ok {
	// 	res.NewValidationErrorResponse(errs).Send(w)
	// 	return
	// }

	entTask, err := t.ent.Task.UpdateOneID(taskID).
		SetTitle(req.Title).
		SetNotes(req.Notes).
		SetIsImportant(req.IsImportant).
		SetIsTimeSenstive(req.IsImportant).
		SetDueAt(req.DueAt).
		SetRemindAt(req.RemindAt).
		// TODO: update at?
		Save(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	task := model.TaskItemResponse{
		Title:          entTask.Title,
		Notes:          entTask.Notes,
		IsTimeSenstive: entTask.IsTimeSenstive,
		IsImportant:    entTask.IsImportant,
		RemindAt:       entTask.RemindAt,
		DueAt:          entTask.DueAt,
	}

	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
}

func (t Task) delete(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = t.ent.Task.DeleteOneID(taskID).Exec(r.Context()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task deleted successfully"))
}
