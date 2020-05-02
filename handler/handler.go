package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

func New() http.Handler {
	r := chi.NewRouter()

	r.Get("/reports", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Report!"))
	})

	return r
}
