package routes

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

    r.Mount("/login", loginRoutes())
    r.Mount("/resource", resourceRoutes())

	return r
}
