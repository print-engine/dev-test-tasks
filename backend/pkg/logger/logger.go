package logger

import (
	"log"
	"os"
	"time"
)

// Logger is a custom logger with structured output
type Logger struct {
	*log.Logger
}

// New creates a new logger instance
func New() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// LogRequest logs HTTP request details
func (l *Logger) LogRequest(method, path string, duration time.Duration, status int) {
	l.Printf("[%s] %s - %d - %v", method, path, status, duration)
}
