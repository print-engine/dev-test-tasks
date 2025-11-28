package main

import (
	"net/http"

	"dev-test-tasks/config"
	"dev-test-tasks/internal/handlers"
	"dev-test-tasks/internal/middleware"
	"dev-test-tasks/pkg/logger"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize logger
	log := logger.New()

	// Create a new mux
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/healthcheck", handlers.HealthCheck)
	mux.HandleFunc("/ws", handlers.WebSocket)

	// Wrap mux with middleware
	handler := middleware.CORS(middleware.Logging(mux, log))

	// Configure server
	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      handler,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
