// Package cards contains the gRPC handler for card generation service.
package cards

import (
	"context"
	"errors"

	"github.com/not-hype-pruduction/bridge-backend/internal/domain/cards"
	"github.com/not-hype-pruduction/bridge-backend/internal/domain/ports/inbound"
	cardsv1 "github.com/not-hype-pruduction/bridge-backend/internal/pb/cards.v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler is the gRPC handler for the card generator service.
type Handler struct {
	cardsv1.UnimplementedCardGeneratorServiceServer
	cardGenerator inbound.CardGenerator
}

// NewHandler creates a new gRPC handler with the given card generator use case.
func NewHandler(cardGenerator inbound.CardGenerator) *Handler {
	return &Handler{
		cardGenerator: cardGenerator,
	}
}

// Register registers the handler with the gRPC server.
func (h *Handler) Register(server *grpc.Server) {
	cardsv1.RegisterCardGeneratorServiceServer(server, h)
}

// GenerateHands handles the GenerateHands gRPC request.
func (h *Handler) GenerateHands(
	ctx context.Context,
	in *cardsv1.GenerateHandsRequest,
) (*cardsv1.GenerateHandsResponse, error) {
	// Validate request
	if err := Validate(in); err != nil {
		return nil, err
	}

	// Map request to domain input
	input := ToGenerateHandsInput(in)

	// Execute use case
	pbn, err := h.cardGenerator.Execute(ctx, input)
	if err != nil {
		if errors.Is(err, cards.ErrImpossibleCardCombination) {
			return nil, status.Error(codes.InvalidArgument, cards.ErrImpossibleCardCombination.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return ToGenerateHandsResponse(pbn), nil
}
