package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/not-hype-pruduction/bridge-backend/internal/app"
	"github.com/not-hype-pruduction/bridge-backend/internal/app/logger"
	"github.com/not-hype-pruduction/bridge-backend/internal/config"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	application := app.New(log, cfg.GPRC.Port, cfg.Random)

	application.GRPCServer.MustRun()

	go func() {
		application.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	log.Info("Gracefully stopped")

}
