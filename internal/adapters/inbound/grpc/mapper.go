// Package grpc contains the gRPC inbound adapter.
package grpc

import (
	"github.com/not-hype-pruduction/bridge-backend/internal/domain/ports/inbound"
	cardsv1 "github.com/not-hype-pruduction/bridge-backend/internal/pb/cards.v1"
)

// ToGenerateHandsInput maps a gRPC GenerateHandsRequest to the domain GenerateHandsInput.
func ToGenerateHandsInput(req *cardsv1.GenerateHandsRequest) inbound.GenerateHandsInput {
	return inbound.GenerateHandsInput{
		MyPointsMin:      req.GetMyPointsMin(),
		MyPointsMax:      req.GetMyPointsMax(),
		PartnerPointsMin: req.GetPartnerPointsMin(),
		PartnerPointsMax: req.GetPartnerPointsMax(),
		Dealer:           req.GetDelaer(),
		North:            int32(req.GetNorth()),
		South:            int32(req.GetSouth()),
		East:             int32(req.GetEast()),
		West:             int32(req.GetWest()),
	}
}

// ToGenerateHandsResponse maps a PBN string to a gRPC GenerateHandsResponse.
func ToGenerateHandsResponse(pbn string) *cardsv1.GenerateHandsResponse {
	return &cardsv1.GenerateHandsResponse{
		Pbn: pbn,
	}
}
