package routes

import (
	"client/internal/middlewares"
	"client/internal/services"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func resourceRoutes() *chi.Mux {
    r := chi.NewRouter()

    r.Use(middlewares.LoggerMiddleware)

    r.Get("/get", getResourceHandler)

    return r
}


func getResourceHandler(w http.ResponseWriter, r *http.Request) {
    authHeader := r.Header.Get("Authorization")

    if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
        http.Error(w, "Missing or invalid Authorization header", http.StatusBadRequest)
        return
    }

    accessToken := strings.TrimPrefix(authHeader, "Bearer ")

    requestBody := services.ResourceRequestBody{
        ClientID: "client-id-example",
    }

    resource, err := services.GetResourceService(requestBody, accessToken)
    if err != nil {
        http.Error(w, "Could not fetch resource", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(resource)
}
