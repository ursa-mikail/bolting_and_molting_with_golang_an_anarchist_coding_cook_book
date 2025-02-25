package logging

import "fmt"

// Logger provides basic logging functionality
type Logger struct{}

// NewLogger creates a new Logger instance
func NewLogger() *Logger {
	return &Logger{}
}

// Error logs an error message
func (l *Logger) Error(message string, err error) {
	fmt.Printf("[ERROR] %s: %v\n", message, err)
}

