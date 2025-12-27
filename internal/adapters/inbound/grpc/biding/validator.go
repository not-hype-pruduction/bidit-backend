// Package biding contains the gRPC handler for biding part of game.
package biding

import bidingv1 "github.com/not-hype-pruduction/bridge-backend/internal/pb/biding.v1"

func Validate(in *bidingv1.MakeBidRequest) error {
	if err := validateBidLevel(in); err != nil {
		return err
	}

	return nil
}

func validateBidLevel(in *bidingv1.MakeBidRequest) error {
	for _, bid := range in.AuctionHistory {
		if bid.Type != *bidingv1.SpecialCall_SPECIAL_CALL_BID.Enum() {
			continue
		}

		if bid.Level < 1 || bid.Level > 7 {
			return ErrInvalidBidLevel
		}
	}

	return nil
}
