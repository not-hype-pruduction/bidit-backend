// Package outbound contains outbound port interfaces for external dependencies.
package outbound

// Logger is the interface for logging operations.
type Logger interface {
	// Debug logs a debug message with optional key-value pairs.
	Debug(msg string, args ...any)
	// Info logs an info message with optional key-value pairs.
	Info(msg string, args ...any)
	// Warn logs a warning message with optional key-value pairs.
	Warn(msg string, args ...any)
	// Error logs an error message with optional key-value pairs.
	Error(msg string, args ...any)
}
