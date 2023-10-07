package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/haileemiu/manage-life/svc/task/model"
)

type Task struct {
}

func New() *Task {
	return &Task{}
}

func (t Task) Routes(r chi.Router) {
	r.Get("/", t.list)
	r.Post("/", t.create)
	r.Get("/{id}", t.getByID)
	r.Put("/{id}", t.update)
	r.Delete("/{id}", t.delete)
}

func (t Task) list(w http.ResponseWriter, r *http.Request) {
	tasks := model.TaskListResponse{
		Tasks: []model.TaskItemResponse{
			{ID: 1, Title: "hard coded 1", Notes: "hard"},
			{ID: 2, Title: "hard coded 2", Notes: "hard"},
		},
	}
	// TODO: handle error -N
	// TODO: set response encoding -N
	json.NewEncoder(w).Encode(tasks)
}

func (t Task) create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Implemented")
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
