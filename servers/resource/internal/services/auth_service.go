package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResourceRequestBody struct {
    ClientID string `json:"client_id"`
}

type AccessTokenRequestBody struct {
    ClientID string `json:"client_id"`
} 

type AccessTokenCheckAccessResp struct {
    Success bool `json:"success"`
} 

func CheckAccessTokenService(requestBody ResourceRequestBody, accessToken string) (AccessTokenCheckAccessResp, error) {
    url := "http://auth-server:8080/access/check-access-token"

    accessCheckBody := AccessTokenRequestBody{
        ClientID: requestBody.ClientID,
    }

    jsonReq, err := json.Marshal(accessCheckBody)
    if err != nil {
        return AccessTokenCheckAccessResp{}, fmt.Errorf("error encoding request body: %w", err)
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
    if err != nil {
        return AccessTokenCheckAccessResp{}, fmt.Errorf("failed to create request: %w", err)
    }

    req.Header.Set("Authorization", "Bearer "+accessToken)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    accessResp, err := client.Do(req)
    if err != nil {
        return AccessTokenCheckAccessResp{}, fmt.Errorf("error sending request to auth service: %w", err)
    }
    defer accessResp.Body.Close()

    if accessResp.StatusCode != http.StatusOK {
        return AccessTokenCheckAccessResp{}, fmt.Errorf("unexpected status code from auth service: %d", accessResp.StatusCode)
    }

    var response AccessTokenCheckAccessResp
    if err := json.NewDecoder(accessResp.Body).Decode(&response); err != nil {
        return AccessTokenCheckAccessResp{}, fmt.Errorf("failed to parse response from auth service: %w", err)
    }

    return response, nil
}
