// Package app contains the application composition and dependency injection.
package app

import (
	"log/slog"

	grpcAdapter "github.com/not-hype-pruduction/bridge-backend/internal/adapters/inbound/grpc"
	grpcadapter "github.com/not-hype-pruduction/bridge-backend/internal/adapters/inbound/grpc"
	bidingHandler "github.com/not-hype-pruduction/bridge-backend/internal/adapters/inbound/grpc/biding"
	cardsHandler "github.com/not-hype-pruduction/bridge-backend/internal/adapters/inbound/grpc/cards"
	loggeradapter "github.com/not-hype-pruduction/bridge-backend/internal/adapters/outbound/logger"
	"github.com/not-hype-pruduction/bridge-backend/internal/application/usecases"
	"github.com/not-hype-pruduction/bridge-backend/internal/domain/biding"
	"github.com/not-hype-pruduction/bridge-backend/internal/domain/biding/systems"
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

	sputnik := systems.SputnikStandard{}

	// Create regestry
	systemsRegestry := biding.NewRegistry(&sputnik)

	// Create use cases
	generateHandsUseCase := usecases.NewGenerateHandsUseCase(loggerAdapter)
	bidingSystems := usecases.NewGetNextBidUseCase(systemsRegestry, loggerAdapter)

	// Create gRPC handler
	cardsHandler := cardsHandler.NewHandler(generateHandsUseCase)
	bidingHandler := bidingHandler.NewHandler(bidingSystems)

	// Create gRPC regestry
	grpcRegistry := grpcAdapter.NewRegistry(
		cardsHandler,
		bidingHandler,
	)

	// Create gRPC server
	grpcServer := grpcadapter.NewServer(log, grpcRegistry, grpcPort)

	return &App{
		GRPCServer: grpcServer,
	}
}
