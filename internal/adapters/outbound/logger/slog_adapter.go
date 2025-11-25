// Package logger contains the slog adapter implementation.
package logger

import (
	"log/slog"

	"github.com/not-hype-pruduction/bridge-backend/internal/domain/ports/outbound"
)

// SlogAdapter adapts the standard slog.Logger to the outbound.Logger interface.
type SlogAdapter struct {
	log *slog.Logger
}

// NewSlogAdapter creates a new SlogAdapter wrapping the provided slog.Logger.
func NewSlogAdapter(log *slog.Logger) outbound.Logger {
	return &SlogAdapter{log: log}
}

// Debug logs a debug message with optional key-value pairs.
func (s *SlogAdapter) Debug(msg string, args ...any) {
	s.log.Debug(msg, args...)
}

// Info logs an info message with optional key-value pairs.
func (s *SlogAdapter) Info(msg string, args ...any) {
	s.log.Info(msg, args...)
}

// Warn logs a warning message with optional key-value pairs.
func (s *SlogAdapter) Warn(msg string, args ...any) {
	s.log.Warn(msg, args...)
}

// Error logs an error message with optional key-value pairs.
func (s *SlogAdapter) Error(msg string, args ...any) {
	s.log.Error(msg, args...)
}
