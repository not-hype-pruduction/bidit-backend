// Package inbound contains inbound port interfaces for the application use cases.
package inbound

import "context"

// GenerateHandsInput contains all parameters for the GenerateHands use case.
type GenerateHandsInput struct {
	MyPointsMin      int32
	MyPointsMax      int32
	PartnerPointsMin int32
	PartnerPointsMax int32
	Dealer           string
	North            int32
	South            int32
	East             int32
	West             int32
}

// CardGenerator is the interface for the card generation use case.
type CardGenerator interface {
	// Execute generates a deal and returns it in PBN format.
	Execute(ctx context.Context, input GenerateHandsInput) (string, error)
}
