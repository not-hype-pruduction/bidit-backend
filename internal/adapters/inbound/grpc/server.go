// Package grpc contains the gRPC inbound adapter.
package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server manages the gRPC server lifecycle.
type Server struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

// NewServer creates a new gRPC server with the provided configuration.
func NewServer(log *slog.Logger, handler *Handler, port int) *Server {
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived, logging.PayloadSent,
		),
	}

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			log.Error("Recovered from panic", slog.Any("panic", p))

			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(InterceptorLogger(log), loggingOpts...),
	))

	handler.Register(gRPCServer)

	return &Server{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// MustRun starts the server and panics on error.
func (s *Server) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

// Run starts the gRPC server.
func (s *Server) Run() error {
	const op = "grpc.Server.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.log.Info("grpc server started", slog.String("addr", l.Addr().String()))

	if err := s.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Stop gracefully stops the gRPC server.
func (s *Server) Stop() {
	const op = "grpc.Server.Stop"

	s.log.With(slog.String("op", op)).
		Info("stopping gRPC server", slog.Int("port", s.port))

	s.gRPCServer.GracefulStop()
}

// InterceptorLogger adapts slog to the grpc middleware logging interface.
func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
