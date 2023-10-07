package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/haileemiu/manage-life/svc/task"
)

// TODO: try chi router nested - see github docs
func main() {
	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second), // TODO: make config variable -N
	)

	taskHdl := task.NewHandler()

	r.Route("/api/tasks", taskHdl.Routes)

	// TODO: make config variable for ip and port -N
	// TODO: add os signal handling for stopping app -N

	http.ListenAndServe(":4000", r)
}
