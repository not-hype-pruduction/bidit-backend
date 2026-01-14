package usecases

import (
	"context"

	"github.com/not-hype-pruduction/bridge-backend/internal/domain/dds"
	"github.com/not-hype-pruduction/bridge-backend/internal/domain/ports/outbound"
)

type GetDDTableUseCase struct {
	solver outbound.DDSolver
}

func (u *GetDDTableUseCase) Execute(ctx context.Context, pbn string) (dds.DDTable, error) {
	return u.solver.CalculateTable(ctx, pbn)
}
