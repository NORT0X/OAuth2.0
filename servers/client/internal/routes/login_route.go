package routes

import (
	"client/internal/middlewares"
	"client/internal/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func loginRoutes() *chi.Mux {
    r := chi.NewRouter()

    r.Use(middlewares.LoggerMiddleware)

    r.Get("/", loginHandler)

    return r
}

var savedAccessToken services.AccessTokenResponse

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	if username == "" {
		http.Error(w, "Missing username", http.StatusBadRequest)
		return
	}

	tokenResponse, err := services.GetAuthTokenService(username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to communicate with authorization service: %v", err), http.StatusInternalServerError)
		return
	}

	requestBody := services.AccessTokenRequestBody{
		AuthorizationCode: tokenResponse.Token,
		ClientId:          "client-id-example",
		AccessTokenSecret: "secret-example",
	}

	accessToken, err := services.GetAccessTokenService(requestBody)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to communicate with access service: %v", err), http.StatusInternalServerError)
		return
	}

    savedAccessToken = accessToken

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": accessToken.Token})
}
