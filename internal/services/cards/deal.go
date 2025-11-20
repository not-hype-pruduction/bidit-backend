package cards

import (
	"errors"
	"strings"
)

type card struct {
	Rank string
	Suit string
}

type hand struct {
	Spades   []card
	Hearts   []card
	Diamonds []card
	Clubs    []card
}

type deal struct {
	My      hand
	Partner hand
	Player1 hand
	Player2 hand
	Dealer  string
}

type pbn struct {
	My      string
	Partner string
	Player1 string
	Player2 string
}

var (
	BADPLAYERPOSITION = errors.New("bad player position, please provide a correct player position")
)

var order = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

func (h *hand) getAmountOfCards() int {
	return len(h.Spades) + len(h.Clubs) + len(h.Diamonds) + len(h.Hearts)
}

func (d *deal) createPBN(north, south, east, west int32) (string, error) {
	var res strings.Builder
	res.WriteString(d.Dealer)
	res.WriteString(":")

	gPBN := pbn{
		My:      handToPBNFormat(&d.My),
		Partner: handToPBNFormat(&d.Partner),
		Player1: handToPBNFormat(&d.Player1),
		Player2: handToPBNFormat(&d.Player2),
	}

	peoples := []string{"N", "E", "S", "W"}

	sIndex := getStartIndex(d.Dealer)

	for i := range 4 {
		switch peoples[i+sIndex%4] {
		case "N":
			str, err := gPBN.getNeedPlayer(north)
			if err != nil {
				return "", err
			}

			res.WriteString(str)

		case "E":
			str, err := gPBN.getNeedPlayer(east)
			if err != nil {
				return "", err
			}

			res.WriteString(str)

		case "S":
			str, err := gPBN.getNeedPlayer(south)
			if err != nil {
				return "", err
			}

			res.WriteString(str)

		case "W":
			str, err := gPBN.getNeedPlayer(west)
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

func (p *pbn) getNeedPlayer(player int32) (string, error) {
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
		return "", BADPLAYERPOSITION
	}
}

func getStartIndex(position string) int {
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

func handToPBNFormat(hand *hand) string {
	var res strings.Builder

	res.WriteString(cardsToPBNFormat(hand.Spades))
	res.WriteString(".")
	res.WriteString(cardsToPBNFormat(hand.Hearts))
	res.WriteString(".")
	res.WriteString(cardsToPBNFormat(hand.Diamonds))
	res.WriteString(".")
	res.WriteString(cardsToPBNFormat(hand.Clubs))

	return res.String()
}

func cardsToPBNFormat(cards []card) string {
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
