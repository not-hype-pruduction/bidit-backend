// Package cards contains the domain entities and business logic for card generation.
package cards

// Card represents a playing card with a rank and suit.
type Card struct {
	Rank string
	Suit string
}

// Hand represents a player's hand of cards organized by suit.
type Hand struct {
	Spades   []Card
	Hearts   []Card
	Diamonds []Card
	Clubs    []Card
}

// Deal represents a complete card deal for all four players.
type Deal struct {
	My      Hand
	Partner Hand
	Player1 Hand
	Player2 Hand
	Dealer  string
}

// PBN represents the PBN (Portable Bridge Notation) format of a deal.
type PBN struct {
	My      string
	Partner string
	Player1 string
	Player2 string
}

// GeneratedCards holds the count of high cards for my hand and partner's hand.
type GeneratedCards struct {
	A, K, Q, J     int8
	A_, K_, Q_, J_ int8
}

// GetAmountOfCards returns the total number of cards in the hand.
func (h *Hand) GetAmountOfCards() int {
	return len(h.Spades) + len(h.Clubs) + len(h.Diamonds) + len(h.Hearts)
}
