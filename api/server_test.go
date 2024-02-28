package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Unit test
func TestHealthCheck(t *testing.T) {
	// Mock request
	req := httptest.NewRequest("GET", "http://localhost.com/healthcheck", nil)
	// Mock response
	w := httptest.NewRecorder()

	// API to test
	healthCheck(w, req, nil)

	res := w.Result()

	// MUST close the body of res
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	// Type of tests you can do
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", res.StatusCode)
	}
	if err != nil {
		t.Errorf("error with HealthCheck")
	}
	expectedRes := "Server is OK"
	if string(body) != expectedRes {
		t.Errorf("Expected %v but got %v instead", expectedRes, string(body))
	}
}
