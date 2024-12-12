package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthTokenResponse struct {
	Token string `json:"token"`
}

func GetAuthTokenService(userID string) (AuthTokenResponse, error) {
	redirectUrl := fmt.Sprintf("http://auth-server:8080/authorization/authorize?user-id=%s", userID)
	authResp, err := http.Get(redirectUrl)
	if err != nil {
		return AuthTokenResponse{}, fmt.Errorf("failed to communicate with authorization service: %w", err)
	}
	defer authResp.Body.Close()

	var tokenResponse AuthTokenResponse
	if err := json.NewDecoder(authResp.Body).Decode(&tokenResponse); err != nil {
		return AuthTokenResponse{}, fmt.Errorf("failed to parse response from authorization service: %w", err)
	}

	return tokenResponse, nil
}
