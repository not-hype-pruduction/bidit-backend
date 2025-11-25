// Package cards contains the domain entities and business logic for card generation.
package cards

import (
	"math/rand"
	"strings"
)

// order defines the ranking order of cards from highest to lowest.
var order = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

// GenerateDeal creates a complete deal from generated cards and a dealer position.
func GenerateDeal(gCards GeneratedCards, dealer string) *Deal {
	var resDeal Deal

	resDeal.Dealer = dealer

	aces := []Card{
		{Rank: "A", Suit: "S"},
		{Rank: "A", Suit: "H"},
		{Rank: "A", Suit: "D"},
		{Rank: "A", Suit: "C"},
	}

	kings := []Card{
		{Rank: "K", Suit: "S"},
		{Rank: "K", Suit: "H"},
		{Rank: "K", Suit: "D"},
		{Rank: "K", Suit: "C"},
	}

	queens := []Card{
		{Rank: "Q", Suit: "S"},
		{Rank: "Q", Suit: "H"},
		{Rank: "Q", Suit: "D"},
		{Rank: "Q", Suit: "C"},
	}

	jacks := []Card{
		{Rank: "J", Suit: "S"},
		{Rank: "J", Suit: "H"},
		{Rank: "J", Suit: "D"},
		{Rank: "J", Suit: "C"},
	}

	var numberCards = []Card{
		// Spades
		{Rank: "2", Suit: "S"},
		{Rank: "3", Suit: "S"},
		{Rank: "4", Suit: "S"},
		{Rank: "5", Suit: "S"},
		{Rank: "6", Suit: "S"},
		{Rank: "7", Suit: "S"},
		{Rank: "8", Suit: "S"},
		{Rank: "9", Suit: "S"},
		{Rank: "T", Suit: "S"},

		// Hearts
		{Rank: "2", Suit: "H"},
		{Rank: "3", Suit: "H"},
		{Rank: "4", Suit: "H"},
		{Rank: "5", Suit: "H"},
		{Rank: "6", Suit: "H"},
		{Rank: "7", Suit: "H"},
		{Rank: "8", Suit: "H"},
		{Rank: "9", Suit: "H"},
		{Rank: "T", Suit: "H"},

		// Diamonds
		{Rank: "2", Suit: "D"},
		{Rank: "3", Suit: "D"},
		{Rank: "4", Suit: "D"},
		{Rank: "5", Suit: "D"},
		{Rank: "6", Suit: "D"},
		{Rank: "7", Suit: "D"},
		{Rank: "8", Suit: "D"},
		{Rank: "9", Suit: "D"},
		{Rank: "T", Suit: "D"},

		// Clubs
		{Rank: "2", Suit: "C"},
		{Rank: "3", Suit: "C"},
		{Rank: "4", Suit: "C"},
		{Rank: "5", Suit: "C"},
		{Rank: "6", Suit: "C"},
		{Rank: "7", Suit: "C"},
		{Rank: "8", Suit: "C"},
		{Rank: "9", Suit: "C"},
		{Rank: "T", Suit: "C"},
	}

	rand.Shuffle(4, func(i, j int) {
		aces[i], aces[j] = aces[j], aces[i]
	})

	rand.Shuffle(4, func(i, j int) {
		kings[i], kings[j] = kings[j], kings[i]
	})

	rand.Shuffle(4, func(i, j int) {
		queens[i], queens[j] = queens[j], queens[i]
	})

	rand.Shuffle(4, func(i, j int) {
		jacks[i], jacks[j] = jacks[j], jacks[i]
	})

	rand.Shuffle(36, func(i, j int) {
		numberCards[i], numberCards[j] = numberCards[j], numberCards[i]
	})

	acesIndex := 0
	kingsIndex := 0
	queensIndex := 0
	jacksIndex := 0
	numbersIndex := 0

	// My cards
	assignCards(&resDeal.My, aces, &acesIndex, int(gCards.A))
	assignCards(&resDeal.My, kings, &kingsIndex, int(gCards.K))
	assignCards(&resDeal.My, queens, &queensIndex, int(gCards.Q))
	assignCards(&resDeal.My, jacks, &jacksIndex, int(gCards.J))

	// Partners cards
	assignCards(&resDeal.Partner, aces, &acesIndex, int(gCards.A_))
	assignCards(&resDeal.Partner, kings, &kingsIndex, int(gCards.K_))
	assignCards(&resDeal.Partner, queens, &queensIndex, int(gCards.Q_))
	assignCards(&resDeal.Partner, jacks, &jacksIndex, int(gCards.J_))

	// Robot1 cards
	if 4-acesIndex != 0 {
		assignCards(&resDeal.Player1, aces, &acesIndex, 4-rand.Intn(4-acesIndex))
	}
	if 4-kingsIndex != 0 {
		assignCards(&resDeal.Player1, kings, &kingsIndex, 4-rand.Intn(4-kingsIndex))
	}

	if 4-queensIndex != 0 {
		assignCards(&resDeal.Player1, queens, &queensIndex, 4-rand.Intn(4-queensIndex))
	}
	if 4-jacksIndex != 0 {
		assignCards(&resDeal.Player1, jacks, &jacksIndex, 4-rand.Intn(4-jacksIndex))
	}

	// Robot2 cards
	if 4-acesIndex != 0 {
		assignCards(&resDeal.Player2, aces, &acesIndex, 4-rand.Intn(4-acesIndex))
	}
	if 4-kingsIndex != 0 {
		assignCards(&resDeal.Player2, kings, &kingsIndex, 4-rand.Intn(4-kingsIndex))
	}

	if 4-queensIndex != 0 {
		assignCards(&resDeal.Player2, queens, &queensIndex, 4-rand.Intn(4-queensIndex))
	}
	if 4-jacksIndex != 0 {
		assignCards(&resDeal.Player2, jacks, &jacksIndex, 4-rand.Intn(4-jacksIndex))
	}

	// Numbers for all players
	assignCards(&resDeal.My, numberCards, &numbersIndex, numbersIndex+13-resDeal.My.GetAmountOfCards())
	assignCards(&resDeal.Partner, numberCards, &numbersIndex, numbersIndex+13-resDeal.Partner.GetAmountOfCards())
	assignCards(&resDeal.Player1, numberCards, &numbersIndex, numbersIndex+13-resDeal.Player1.GetAmountOfCards())
	assignCards(&resDeal.Player2, numberCards, &numbersIndex, numbersIndex+13-resDeal.Player2.GetAmountOfCards())

	return &resDeal
}

// assignCards distributes cards from a deck to a hand.
func assignCards(gHand *Hand, cards []Card, index *int, count int) {
	for ; *index < count && *index < len(cards); *index++ {
		tmp := cards[*index]

		switch tmp.Suit {
		case "S":
			gHand.Spades = append(gHand.Spades, tmp)

		case "H":
			gHand.Hearts = append(gHand.Hearts, tmp)

		case "D":
			gHand.Diamonds = append(gHand.Diamonds, tmp)

		case "C":
			gHand.Clubs = append(gHand.Clubs, tmp)
		}
	}
}

// CreatePBN generates a PBN (Portable Bridge Notation) string from a deal.
func (d *Deal) CreatePBN(north, south, east, west int32) (string, error) {
	var res strings.Builder
	res.WriteString(d.Dealer)
	res.WriteString(":")

	gPBN := PBN{
		My:      HandToPBNFormat(&d.My),
		Partner: HandToPBNFormat(&d.Partner),
		Player1: HandToPBNFormat(&d.Player1),
		Player2: HandToPBNFormat(&d.Player2),
	}

	peoples := []string{"N", "E", "S", "W"}

	sIndex := GetStartIndex(d.Dealer)

	for i := range 4 {
		switch peoples[(i+sIndex)%4] {
		case "N":
			str, err := gPBN.GetNeedPlayer(north)
			if err != nil {
				return "", err
			}

			res.WriteString(str)

		case "E":
			str, err := gPBN.GetNeedPlayer(east)
			if err != nil {
				return "", err
			}

			res.WriteString(str)

		case "S":
			str, err := gPBN.GetNeedPlayer(south)
			if err != nil {
				return "", err
			}

			res.WriteString(str)

		case "W":
			str, err := gPBN.GetNeedPlayer(west)
			if err != nil {
				return "", err
			}

			res.WriteString(str)
		}

		if i < 3 {
			res.WriteString(" ")
		}
	}

	return res.String(), nil
}

// GetNeedPlayer returns the PBN string for a specific player position.
func (p *PBN) GetNeedPlayer(player int32) (string, error) {
	switch player {
	case 1:
		return p.My, nil

	case 2:
		return p.Partner, nil

	case 3:
		return p.Player1, nil

	case 4:
		return p.Player2, nil

	default:
		return "", ErrBadPlayerPosition
	}
}

// GetStartIndex returns the starting index based on dealer position.
func GetStartIndex(position string) int {
	switch position {
	case "N":
		return 0
	case "E":
		return 1
	case "S":
		return 2
	case "W":
		return 3
	}

	return 0
}

// HandToPBNFormat converts a hand to PBN format string.
func HandToPBNFormat(hand *Hand) string {
	var res strings.Builder

	res.WriteString(CardsToPBNFormat(hand.Spades))
	res.WriteString(".")
	res.WriteString(CardsToPBNFormat(hand.Hearts))
	res.WriteString(".")
	res.WriteString(CardsToPBNFormat(hand.Diamonds))
	res.WriteString(".")
	res.WriteString(CardsToPBNFormat(hand.Clubs))

	return res.String()
}

// CardsToPBNFormat converts a slice of cards to PBN format string.
func CardsToPBNFormat(cards []Card) string {
	var res strings.Builder

	for _, rank := range order {
		for _, c := range cards {
			if c.Rank == rank {
				res.WriteString(c.Rank)
			}
		}
	}

	return res.String()
}
