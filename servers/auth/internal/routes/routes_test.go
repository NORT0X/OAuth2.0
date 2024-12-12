package routes

import (
	"auth/internal/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorizationHandler_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/authorization/authorize?user-id=testuser", nil)
	rec := httptest.NewRecorder()

	authorizationHandler(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "Expected status OK")
	assert.Contains(t, rec.Body.String(), "token", "Response should contain a token")

	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	token := response["token"]
	claims, err := utils.ValidateAuthorizationToken(token)
	assert.NoError(t, err, "Token should be valid")
	assert.Equal(t, "testuser", claims.Username, "Token claims should contain the correct username")
}

func TestAuthorizationHandler_MissingUserID(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/authorization/authorize", nil)
	rec := httptest.NewRecorder()

	authorizationHandler(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code, "Expected status Internal Server Error")
	assert.Contains(t, rec.Body.String(), "Missing user-id query parameter", "Response should indicate failure to generate token")
}


func TestAccessTokenHandler_Success(t *testing.T) {
    userId := "testuser"
    authToken, err := utils.GenerateAuthorizationToken(userId)
    if err != nil {
        t.Fatalf("Failed to generate authorization token: %v", err)
    }

    requestBody := `{
        "authorization_code": "` + authToken + `",
        "client_id": "testclient",
        "access_secret": "secret"
    }`
    req := httptest.NewRequest(http.MethodPost, "/access/access-token", bytes.NewBufferString(requestBody))
    req.Header.Set("Content-Type", "application/json")
    rec := httptest.NewRecorder()

    accessTokenHandler(rec, req)

    assert.Equal(t, http.StatusOK, rec.Code, "Expected status OK")
    assert.Contains(t, rec.Body.String(), "token", "Response should contain an access token")

    var response map[string]string
    err = json.Unmarshal(rec.Body.Bytes(), &response)
    assert.NoError(t, err, "Response should be valid JSON")

    accessToken := response["token"]
    _, err = utils.ValidateAccessToken(accessToken)
    assert.NoError(t, err, "Access token should be valid")
}

func TestAccessTokenHandler_InvalidAuthCode(t *testing.T) {
	requestBody := `{
		"authorization_code": "invalidAuthCode",
		"client_id": "testclient",
		"access_secret": "secret"
	}`
	req := httptest.NewRequest(http.MethodPost, "/access/access-token", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	accessTokenHandler(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code, "Expected status Bad Request")
	assert.Contains(t, rec.Body.String(), "Invalid authorization code", "Response should indicate invalid authorization code")
}

