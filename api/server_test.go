package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Unit test
func TestHealthCheck(t *testing.T) {
	// Mock request
	req := httptest.NewRequest("GET", "http:localhost.com", nil)
	// Mock response
	w := httptest.NewRecorder()

	// API to test
	healthCheck(w, req)

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

func TestShortenUrl(t *testing.T) {
	// Mock request
	reqBody := strings.NewReader(`{"uri": "testing"}`)
	req := httptest.NewRequest("GET", "http:localhost.com", reqBody)
	// Mock response
	w := httptest.NewRecorder()

	// API to test
	shortenUrl(w, req)

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

	t.Logf("%v", body)
}
