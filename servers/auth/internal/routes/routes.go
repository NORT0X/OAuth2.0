package routes

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/authorization", authRoutes())
    r.Mount("/access", accessRoutes())

	return r
}
