package biding

import "context"

// BiddingSystem - it is interface for every bidding system (SAYC, 2/1 and ...)
type BiddingSystem interface {
	Name() string
	GetBid(ctx context.Context, in BidState) (Call, error)
}
