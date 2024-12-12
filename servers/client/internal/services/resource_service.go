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

type ResourceResponse struct {
    Resource string `json:"resource"`
} 

func GetResourceService(requestBody ResourceRequestBody, accessToken string) (ResourceResponse, error) {
    url := "http://resource-server:8082/resource/get"

    resourceJsonReq, err := json.Marshal(requestBody)
    if err != nil {
        return ResourceResponse{}, fmt.Errorf("error encoding request body: %w", err)
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(resourceJsonReq))
    if err != nil {
        return ResourceResponse{}, fmt.Errorf("failed to create request: %w", err)
    }

    req.Header.Set("Authorization", "Bearer "+accessToken)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resourceResp, err := client.Do(req)
    if err != nil {
        return ResourceResponse{}, fmt.Errorf("failed to send request to resource service: %w", err)
    }
    defer resourceResp.Body.Close()

    if resourceResp.StatusCode != http.StatusOK {
        return ResourceResponse{}, fmt.Errorf("unexpected status code from resource service: %d", resourceResp.StatusCode)
    }

    var response ResourceResponse
    if err := json.NewDecoder(resourceResp.Body).Decode(&response); err != nil {
        return ResourceResponse{}, fmt.Errorf("failed to parse response from resource service: %w", err)
    }

    return response, nil
}
