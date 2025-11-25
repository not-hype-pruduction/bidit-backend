// Package app contains the application composition and dependency injection.
package app

import (
	"log/slog"

	grpcadapter "github.com/not-hype-pruduction/bridge-backend/internal/adapters/inbound/grpc"
	loggeradapter "github.com/not-hype-pruduction/bridge-backend/internal/adapters/outbound/logger"
	"github.com/not-hype-pruduction/bridge-backend/internal/application/usecases"
)

// App represents the main application with all its dependencies.
type App struct {
	GRPCServer *grpcadapter.Server
}

// New creates a new application with all dependencies wired together.
func New(
	log *slog.Logger,
	grpcPort int,
) *App {
	// Create logger adapter
	loggerAdapter := loggeradapter.NewSlogAdapter(log)

	// Create use cases
	generateHandsUseCase := usecases.NewGenerateHandsUseCase(loggerAdapter)

	// Create gRPC handler
	grpcHandler := grpcadapter.NewHandler(generateHandsUseCase)

	// Create gRPC server
	grpcServer := grpcadapter.NewServer(log, grpcHandler, grpcPort)

	return &App{
		GRPCServer: grpcServer,
	}
}
