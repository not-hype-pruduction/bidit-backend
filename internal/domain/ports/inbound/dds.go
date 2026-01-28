package inbound

import (
	"context"

	"github.com/not-hype-pruduction/bridge-backend/internal/domain/dds"
)

type DDS interface {
	Execute(ctx context.Context, pbn string) (dds.DDTable, error)
}
