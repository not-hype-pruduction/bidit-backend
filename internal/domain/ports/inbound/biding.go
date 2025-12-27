package inbound

import (
	"context"

	"github.com/not-hype-pruduction/bridge-backend/internal/domain/biding"
)

type Biding interface {
	Execute(ctx context.Context, state biding.BidState) (biding.Call, error)
}
