package usecases

import (
	"context"

	"github.com/not-hype-pruduction/bridge-backend/internal/domain/biding"
	"github.com/not-hype-pruduction/bridge-backend/internal/domain/ports/outbound"
	"github.com/not-hype-pruduction/bridge-backend/internal/lib/logger/sl"
)

type GetNextBidUseCase struct {
	registry *biding.Registry
	log      outbound.Logger
}

func NewGetNextBidUseCase(r *biding.Registry, log outbound.Logger) *GetNextBidUseCase {
	return &GetNextBidUseCase{
		registry: r,
		log:      log,
	}
}

func (uc *GetNextBidUseCase) Execute(ctx context.Context, state biding.BidState) (biding.Call, error) {
	system, err := uc.registry.Get(state.SystemName)
	if err != nil {
		uc.log.Error("invalid system name", sl.Err(err))
		return biding.Call{}, err
	}
	return system.GetBid(ctx, state)
}
