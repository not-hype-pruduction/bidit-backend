// Package usecases contains the application use cases.
package usecases

import (
	"context"

	"github.com/not-hype-pruduction/bridge-backend/internal/domain/cards"
	"github.com/not-hype-pruduction/bridge-backend/internal/domain/ports/inbound"
	"github.com/not-hype-pruduction/bridge-backend/internal/domain/ports/outbound"
)

// GenerateHandsUseCase implements the CardGenerator interface for generating card hands.
type GenerateHandsUseCase struct {
	log outbound.Logger
}

// NewGenerateHandsUseCase creates a new instance of GenerateHandsUseCase.
func NewGenerateHandsUseCase(log outbound.Logger) *GenerateHandsUseCase {
	return &GenerateHandsUseCase{
		log: log,
	}
}

// Execute generates card hands based on the input parameters and returns PBN string.
func (uc *GenerateHandsUseCase) Execute(ctx context.Context, input inbound.GenerateHandsInput) (string, error) {
	generatedCards, err := cards.GenerateCardsWithPoints(
		input.MyPointsMin, input.MyPointsMax,
		input.PartnerPointsMin, input.PartnerPointsMax,
	)
	if err != nil {
		return "", err
	}

	deal := cards.GenerateDeal(*generatedCards, input.Dealer)

	return deal.CreatePBN(input.North, input.South, input.East, input.West)
}
