package routes

import (
	"auth/internal/middlewares"
	"auth/internal/services"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)


func accessRoutes() *chi.Mux {
	r := chi.NewRouter()

    r.Use(middlewares.LoggerMiddleware)

	r.Post("/access-token", accessTokenHandler)
	r.Post("/check-access-token", checkAccessTokenHandler)

	return r
}

type AccessTokenRequestBody struct {
	AuthorizationCode string `json:"authorization_code"`
	ClientId          string `json:"client_id"`
	AccessTokenSecret string `json:"access_secret"`
}

func accessTokenHandler(w http.ResponseWriter, r *http.Request) {
	var body AccessTokenRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_, err := services.ValidateAuthorizationToken(body.AuthorizationCode)

	if err != nil {
		http.Error(w, "Invalid authorization code", http.StatusBadRequest)
		return
	}

	accessToken, err := services.GenerateAccessToken(body.ClientId, body.AccessTokenSecret)

	if err != nil {
		http.Error(w, "Could not generate access token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": accessToken})

}

type AccessTokenCheckBody struct {
    ClientID string `json:"client_id"`
}

func checkAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
    authHeader := r.Header.Get("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Missing or invalid Authorization header", http.StatusBadRequest)
		return
	}

	accessToken := strings.TrimPrefix(authHeader, "Bearer ")

	var body AccessTokenCheckBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	_, err := services.ValidateAccessToken(accessToken, body.ClientID)

	if err != nil {
		http.Error(w, "Invalid access token", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}
