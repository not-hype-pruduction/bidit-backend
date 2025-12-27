package biding

import (
	"github.com/not-hype-pruduction/bridge-backend/internal/domain/biding"
	bidingv1 "github.com/not-hype-pruduction/bridge-backend/internal/pb/biding.v1"
)

func ToBidState(in *bidingv1.MakeBidRequest) biding.BidState {
	var auctionHistory []biding.Call

	for _, i := range in.AuctionHistory {
		auctionHistory = append(
			auctionHistory,
			biding.Call{
				Level: int(i.Level),
				Suit:  biding.Suit(i.Suit.Number()),
				Type:  i.Type.String(),
			},
		)
	}

	return biding.BidState{
		Hand:           in.PlayerHand.Pbn,
		AuctionHistory: auctionHistory,
		Vulnerability:  in.Vulnerability.String(),
		SystemName:     in.SystemName.String(),
	}
}

func ToMakeBidResponse(bid *biding.Call) *bidingv1.MakeBidResponse {
	bidType := 0

	switch bid.Type {
	case "PASS":
		bidType = 1
	case "DOUBLE":
		bidType = 2
	case "REDOUBLE":
		bidType = 3
	case "BID":
		bidType = 4
	}

	return &bidingv1.MakeBidResponse{
		NextCall: &bidingv1.Call{
			Type:        bidingv1.SpecialCall(bidType),
			Level:       int32(bid.Level),
			Suit:        bidingv1.Suit(bid.Suit),
			Explanation: "heheheheheh",
		},
	}
}
