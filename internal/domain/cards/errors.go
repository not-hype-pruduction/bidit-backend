// Package cards contains the domain entities and business logic for card generation.
package cards

import "errors"

// Domain-specific errors for card generation.
var (
	// ErrImpossibleCardCombination is returned when the selected card combination cannot occur.
	ErrImpossibleCardCombination = errors.New("the selected card combination cannot occur")

	// ErrBadPlayerPosition is returned when an invalid player position is provided.
	ErrBadPlayerPosition = errors.New("bad player position, please provide a correct player position")

	// ErrPointsInvalid is returned when the amount of points is incorrect.
	ErrPointsInvalid = errors.New("amount of points is incorrect, please provide a correct amount of points")

	// ErrDealerInvalid is returned when the dealer is incorrect.
	ErrDealerInvalid = errors.New("the dealer is incorrect, please provide a correct dealer")
)
