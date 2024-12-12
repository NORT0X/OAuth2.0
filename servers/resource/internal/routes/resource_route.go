package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"resource/internal/middlewares"
	"resource/internal/services"
	"strings"

	"github.com/go-chi/chi/v5"
)

func resourceRoutes() *chi.Mux {
    r := chi.NewRouter()

    r.Use(middlewares.LoggerMiddleware)

    r.Post("/get", getResourceHandler)

    return r
}

func getResourceHandler(w http.ResponseWriter, r *http.Request) {
    authHeader := r.Header.Get("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Missing or invalid Authorization header", http.StatusBadRequest)
		return
	}

	accessToken := strings.TrimPrefix(authHeader, "Bearer ")


    var body services.ResourceRequestBody 
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    accessResponse, err := services.CheckAccessTokenService(body, accessToken)

    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to communicate with auth service: %v", err), http.StatusBadRequest)
        return
    }

    if accessResponse.Success == false {
        http.Error(w, fmt.Sprintf("Invalid access token"), http.StatusNonAuthoritativeInfo)
        return
    }


    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"resource": "Some important resource"})
}
