package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/haileemiu/manage-life/pkg/client"
	"github.com/haileemiu/manage-life/svc/task"
)

func runAPI() (err error) {
	// Setup service dependencies
	entClient, err := client.GetEnt()
	if err != nil {
		return
	}
	defer entClient.Close() // TODO: add to HTTP close handler -N

	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second), // TODO: make config variable -N
	)

	taskHdl := task.NewHandler(entClient)

	r.Route("/api/tasks", taskHdl.Routes)

	// TODO: make config variable for ip and port -N
	// TODO: add os signal handling for stopping app -N

	return http.ListenAndServe(":4000", r)
}

// TODO: try chi router nested - see github docs
func main() {
	if err := runAPI(); err != nil {
		fmt.Fprintf(os.Stderr, "error running api: %s\n", err)
		os.Exit(1)
	}
}
