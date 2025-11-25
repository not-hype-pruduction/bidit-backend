// Package cards contains the gRPC handler for card generation service.
package cards

import (
	"github.com/not-hype-pruduction/bridge-backend/internal/domain/cards"
	cardsv1 "github.com/not-hype-pruduction/bridge-backend/internal/pb/cards.v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Validate validates the GenerateHandsRequest.
func Validate(in *cardsv1.GenerateHandsRequest) error {
	if err := validatePoints(in); err != nil {
		return err
	}

	if err := validateDealer(in.GetDelaer()); err != nil {
		return err
	}

	return nil
}

// validatePoints validates the points range in the request.
func validatePoints(in *cardsv1.GenerateHandsRequest) error {
	if in.MyPointsMin > in.MyPointsMax {
		return status.Error(codes.InvalidArgument, "my_points_min cannot be greater than my_points_max")
	}

	if in.PartnerPointsMin > in.PartnerPointsMax {
		return status.Error(codes.InvalidArgument, "partner_points_min cannot be greater than partner_points_max")
	}

	if in.MyPointsMax+in.PartnerPointsMax > 40 {
		return status.Error(codes.InvalidArgument, cards.ErrPointsInvalid.Error())
	}

	if in.MyPointsMin < 0 || in.MyPointsMax < 0 || in.PartnerPointsMin < 0 || in.PartnerPointsMax < 0 {
		return status.Error(codes.InvalidArgument, "points cannot be negative")
	}

	return nil
}

// validateDealer validates the dealer position.
func validateDealer(dealer string) error {
	switch dealer {
	case "S", "W", "N", "E":
		return nil
	default:
		return status.Error(codes.InvalidArgument, cards.ErrDealerInvalid.Error())
	}
}
