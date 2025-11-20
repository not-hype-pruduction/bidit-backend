package app

import (
	"log/slog"

	grpcapp "github.com/not-hype-pruduction/bridge-backend/internal/app/grpc"
	"github.com/not-hype-pruduction/bridge-backend/internal/services/cards"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	random bool,
) *App {

	cardsService := cards.New(log, random)

	grpcApp := grpcapp.New(log, &cardsService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
