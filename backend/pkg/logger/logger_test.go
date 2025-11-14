package logger

import (
	"testing"
)

func TestNew(t *testing.T) {
	log := New()
	if log == nil {
		t.Fatal("New() returned nil logger")
	}
	if log.Logger == nil {
		t.Fatal("New() returned logger with nil Logger field")
	}
}

func TestLogRequest(t *testing.T) {
	// Create a logger
	log := New()

	// This test just ensures LogRequest doesn't panic
	// In a real scenario, you might want to capture and verify the output
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("LogRequest panicked: %v", r)
		}
	}()

	// Log a request - should not panic
	log.LogRequest("GET", "/test", 100, 200)
}
