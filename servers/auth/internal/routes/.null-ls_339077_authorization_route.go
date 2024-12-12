package routes

import (
	"auth/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var authCodes = make(map[string]string)

func authRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/login", loginHandler)
	r.Get("/authorize", authorizationHandler)

	return r
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	if username == "" {
		http.Error(w, "Missing username", http.StatusBadRequest)
		return
	}

	authCode := fmt.Sprintf("%s-authcode", username)
	authCodes[authCode] = username

	http.Redirect(w, r, "/auth", http.StatusFound)
}

func authorizationHandler(w http.ResponseWriter, r *http.Request) {
	authCode := r.FormValue("code")
	username, exists := authCodes[authCode]

	if !exists {
		http.Error(w, "Invalid auth code", http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateToken(username)

	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

type AccessTokenRequestBody struct {
    AuthorizationCode string `json:"authorization_code"`
    ClientId string `json:"client_id"`
    AccessTokenSecret string `json:"access_secret"`

}

func accessTokenHandler(w http.ResponseWriter, r *http.Request) {
    var body AccessTokenRequestBody
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()
    


}

