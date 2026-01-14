package outbound

import (
	"context"

	"github.com/not-hype-pruduction/bridge-backend/internal/domain/dds"
)

type DDSolver interface {
	CalculateTable(ctx context.Context, pbn string) (dds.DDTable, error)
}
