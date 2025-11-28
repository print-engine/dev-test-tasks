package models

import "time"

// HealthCheckResponse represents the health check response
type HealthCheckResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}
