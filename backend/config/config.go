package config

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Server ServerConfig
	DB     DatabaseConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	Path string // Path to SQLite database file
}

// Load reads configuration from environment variables
// It first attempts to load a .env file if present
func Load() *Config {
	// Try to load .env file from current directory
	// Silently ignore if file doesn't exist (not an error)
	if err := godotenv.Load(); err != nil {
		// Try loading from config directory
		if err := godotenv.Load("config/.env"); err != nil {
			// .env file is optional, so we just log a debug message
			log.Println("No .env file found, using environment variables and defaults")
		}
	}

	return &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "8080"),
			ReadTimeout:  getDurationEnv("READ_TIMEOUT", 10*time.Second),
			WriteTimeout: getDurationEnv("WRITE_TIMEOUT", 10*time.Second),
			IdleTimeout:  getDurationEnv("IDLE_TIMEOUT", 60*time.Second),
		},
		DB: DatabaseConfig{
			Path: getEnv("DB_PATH", getDefaultDBPath()),
		},
	}
}

// getDefaultDBPath returns the default database path
func getDefaultDBPath() string {
	return filepath.Join("data", "app.db")
}

// getEnv reads an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getDurationEnv reads a duration environment variable or returns a default value
func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		return defaultValue
	}
	return duration
}
