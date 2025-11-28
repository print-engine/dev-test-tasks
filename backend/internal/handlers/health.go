package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"dev-test-tasks/internal/models"
)

// HealthCheck handles the /healthcheck endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := models.HealthCheckResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic("Failed to encode response")
	}
}
