package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AccessTokenResponse struct {
    Token string `json:"token"`
}

type AccessTokenRequestBody struct {
	AuthorizationCode string `json:"authorization_code"`
	ClientId          string `json:"client_id"`
	AccessTokenSecret string `json:"access_secret"`
}

func GetAccessTokenService(requestBody AccessTokenRequestBody) (AccessTokenResponse, error) {
	accessUrl := "http://auth-server:8080/access/access-token"

	accessJsonReq, err := json.Marshal(requestBody)
	if err != nil {
		return AccessTokenResponse{}, fmt.Errorf("error encoding request body: %w", err)
	}

	accessResp, err := http.Post(accessUrl, "application/json", bytes.NewBuffer(accessJsonReq))
	if err != nil {
		return AccessTokenResponse{}, fmt.Errorf("error sending POST request to access service: %w", err)
	}
	defer accessResp.Body.Close()

	if accessResp.StatusCode != http.StatusOK {
		return AccessTokenResponse{}, fmt.Errorf("unexpected status code from access service: %d", accessResp.StatusCode)
	}

	var response AccessTokenResponse
	if err := json.NewDecoder(accessResp.Body).Decode(&response); err != nil {
		return AccessTokenResponse{}, fmt.Errorf("failed to parse response from access service: %w", err)
	}

	return response, nil
}
