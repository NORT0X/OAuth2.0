package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLoginHandler_Success tests the login handler when a valid username is provided
func TestLoginHandler_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/login?username=testuser", nil)
	rec := httptest.NewRecorder()

	loginHandler(rec, req)

	// Check if the response status is OK
	assert.Equal(t, http.StatusOK, rec.Code, "Expected status OK")
	assert.Contains(t, rec.Body.String(), "token", "Response should contain a token")

	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	token := response["token"]
	assert.NotEmpty(t, token, "Token should not be empty")
}

// TestLoginHandler_MissingUsername tests the login handler when the username is missing
func TestLoginHandler_MissingUsername(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/login", nil)
	rec := httptest.NewRecorder()

	loginHandler(rec, req)

	// Check if the response status is Bad Request
	assert.Equal(t, http.StatusBadRequest, rec.Code, "Expected status Bad Request")
	assert.Contains(t, rec.Body.String(), "Missing username", "Response should indicate missing username")
}

// TestResourceHandler_Success tests the resource handler when a valid token is provided
func TestResourceHandler_Success(t *testing.T) {
	// Simulate a request with a valid token
	req := httptest.NewRequest(http.MethodGet, "/resource/get?token=access-token-mock", nil)
	rec := httptest.NewRecorder()

	getResourceHandler(rec, req)

	// Check if the response status is OK
	assert.Equal(t, http.StatusOK, rec.Code, "Expected status OK")
	assert.Contains(t, rec.Body.String(), "resource", "Response should contain a resource")

	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	resource := response["resource"]
	assert.Equal(t, "mock-resource", resource, "Resource should match mock resource")
}

// TestResourceHandler_Failure tests the resource handler when an invalid token is provided
func TestResourceHandler_Failure(t *testing.T) {
	// Simulate a request with an invalid token
	req := httptest.NewRequest(http.MethodGet, "/resource/get?token=invalid-token", nil)
	rec := httptest.NewRecorder()

	getResourceHandler(rec, req)

	// Check if the response status is Not Found
	assert.Equal(t, http.StatusNotFound, rec.Code, "Expected status Not Found")
	assert.Contains(t, rec.Body.String(), "Could not fetch resource", "Response should indicate failure to fetch resource")
}

