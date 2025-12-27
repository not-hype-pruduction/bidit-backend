// Package biding contains the gRPC handler for biding part of game.
package biding

import (
	"context"

	"github.com/not-hype-pruduction/bridge-backend/internal/domain/ports/inbound"
	bidingv1 "github.com/not-hype-pruduction/bridge-backend/internal/pb/biding.v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler is the gRPC handler for the biding part of game.
type Handler struct {
	bidingv1.UnimplementedBidingServiceServer
	biding inbound.Biding
}

// NewHandler creates a new gRPC handler with the given biding use case.
func NewHandler(biding inbound.Biding) *Handler {
	return &Handler{
		biding: biding,
	}
}

// Register registers the handler with the gRPC server.
func (h *Handler) Register(server *grpc.Server) {
	bidingv1.RegisterBidingServiceServer(server, h)
}

func (h *Handler) MakeBid(
	ctx context.Context,
	in *bidingv1.MakeBidRequest,
) (*bidingv1.MakeBidResponse, error) {
	err := Validate(in)
	if err != nil {
		return nil, err
	}

	bid, err := h.biding.Execute(ctx, ToBidState(in))
	if err != nil {
		return nil, status.Error(codes.Aborted, "the server cannot make bid")
	}

	return ToMakeBidResponse(&bid), nil
}
