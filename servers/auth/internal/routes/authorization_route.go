package routes

import (
	"auth/internal/middlewares"
	"auth/internal/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)


func authRoutes() *chi.Mux {
	r := chi.NewRouter()

    r.Use(middlewares.LoggerMiddleware)

	r.Get("/authorize", authorizationHandler)

	return r
}

func authorizationHandler(w http.ResponseWriter, r *http.Request) {
    userId := r.FormValue("user-id")
    if userId == "" {
        http.Error(w, "Missing user-id query parameter", http.StatusInternalServerError)
        return
    }

	token, err := services.GenerateAuthorizationToken(userId)

	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
