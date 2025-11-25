// Package app contains the application composition and dependency injection.
package app

import (
	"log/slog"

	grpcAdapter "github.com/not-hype-pruduction/bridge-backend/internal/adapters/inbound/grpc"
	grpcadapter "github.com/not-hype-pruduction/bridge-backend/internal/adapters/inbound/grpc"
	cardsHandler "github.com/not-hype-pruduction/bridge-backend/internal/adapters/inbound/grpc/cards"
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
	cardsHandler := cardsHandler.NewHandler(generateHandsUseCase)

	// Create gRPC regestry
	grpcRegistry := grpcAdapter.NewRegistry(
		cardsHandler,
	)

	// Create gRPC server
	grpcServer := grpcadapter.NewServer(log, grpcRegistry, grpcPort)

	return &App{
		GRPCServer: grpcServer,
	}
}
