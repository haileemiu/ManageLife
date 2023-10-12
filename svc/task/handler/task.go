package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/haileemiu/manage-life/ent"
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
	// TODO: handle error -N
	// TODO: set response encoding -N
	// TODO: convert from ent model to task model -N
	json.NewEncoder(w).Encode(tasks)
}

func (t Task) create(w http.ResponseWriter, r *http.Request) {
	req := model.TaskCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if ok, errs := req.Validate(); !ok {
		res.NewValidationErrorResponse(errs).Send(w)
		return
	}

	task, err := t.ent.Task.Create().
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

	// TODO: handle error
	// TODO: set response encoding
	// TODO: convert from ent model to task model
	json.NewEncoder(w).Encode(task)
}

func (t Task) getByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Implemented")
}

func (t Task) update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Implemented")
}

func (t Task) delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Implemented")
}