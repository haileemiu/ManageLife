package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// TODO: try chi router nested - see github docs
func main() {
	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second), // TODO: make config variable
	)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		test := r.URL.Query().Get("test")
		if test == "" {
			test = "World"
		}

		fmt.Fprintf(w, "Hi %s", test)
	})

	http.ListenAndServe(":4000", r)
}
