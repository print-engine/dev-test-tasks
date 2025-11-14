package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthCheckHandler(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthCheckHandler)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the Content-Type header
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v",
			contentType, "application/json")
	}

	// Parse the response body
	var response HealthCheckResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	// Check the response fields
	if response.Status != "ok" {
		t.Errorf("handler returned wrong status: got %v want %v",
			response.Status, "ok")
	}

	// Check that timestamp is recent (within last second)
	if time.Since(response.Timestamp) > time.Second {
		t.Errorf("timestamp is not recent: %v", response.Timestamp)
	}
}

func TestLoggingMiddleware(t *testing.T) {
	logger := NewLogger()

	// Create a simple test handler
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test response"))
	})

	// Wrap with logging middleware
	handler := loggingMiddleware(testHandler, logger)

	// Create a test request
	req := httptest.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body := rr.Body.String()
	if body != "test response" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, "test response")
	}
}

func TestResponseWriter(t *testing.T) {
	// Create a test ResponseWriter
	rr := httptest.NewRecorder()
	wrapped := &responseWriter{ResponseWriter: rr, statusCode: http.StatusOK}

	// Test WriteHeader
	wrapped.WriteHeader(http.StatusNotFound)

	if wrapped.statusCode != http.StatusNotFound {
		t.Errorf("statusCode not captured: got %v want %v",
			wrapped.statusCode, http.StatusNotFound)
	}

	if rr.Code != http.StatusNotFound {
		t.Errorf("status code not written: got %v want %v",
			rr.Code, http.StatusNotFound)
	}
}
