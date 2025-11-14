package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"dev-test-tasks/pkg/logger"
)

func TestLogging(t *testing.T) {
	log := logger.New()

	// Create a simple test handler
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test response"))
	})

	// Wrap with logging middleware
	handler := Logging(testHandler, log)

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
