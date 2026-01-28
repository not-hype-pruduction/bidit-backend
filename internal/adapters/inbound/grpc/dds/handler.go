package dds

import (
	"context"

	"github.com/not-hype-pruduction/bridge-backend/internal/domain/ports/inbound"
	ddsv1 "github.com/not-hype-pruduction/bridge-backend/internal/pb/dds.v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler is the gRPC handler for the DDS service.
type Handler struct {
	ddsv1.UnimplementedDDSServiceServer
	dds inbound.DDS
}

// NewHandler creates a new gRPC handler with the given DDS use case.
func NewHandler(dds inbound.DDS) *Handler {
	return &Handler{
		dds: dds,
	}
}

// Register registers the handler with the gRPC server.
func (h *Handler) Register(server *grpc.Server) {
	ddsv1.RegisterDDSServiceServer(server, h)
}

func (h *Handler) GetDDTable(
	ctx context.Context,
	in *ddsv1.GetDDTableRequest,
) (*ddsv1.GetDDTableResponse, error) {
	err := validatePBN(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ddTable, err := h.dds.Execute(ctx, in.Pbn)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return ToGetDDTableResponse(&ddTable), nil
}
