package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

// Logger is a custom logger with structured output
type Logger struct {
	*log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// LogRequest logs HTTP request details
func (l *Logger) LogRequest(method, path string, duration time.Duration, status int) {
	l.Printf("[%s] %s - %d - %v", method, path, status, duration)
}

// loggingMiddleware wraps an http.Handler and logs request details
func loggingMiddleware(next http.Handler, logger *Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a response writer wrapper to capture status code
		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Call the next handler
		next.ServeHTTP(wrapped, r)

		// Log the request
		duration := time.Since(start)
		logger.LogRequest(r.Method, r.URL.Path, duration, wrapped.statusCode)
	})
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// HealthCheckResponse represents the health check response
type HealthCheckResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

// healthCheckHandler handles the /healthcheck endpoint
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthCheckResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	logger := NewLogger()

	// Create a new mux
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/healthcheck", healthCheckHandler)

	// Wrap mux with logging middleware
	handler := loggingMiddleware(mux, logger)

	// Configure server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Printf("Server starting on port %s", port)
	if err := server.ListenAndServe(); err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
