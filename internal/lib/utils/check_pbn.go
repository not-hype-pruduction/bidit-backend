package utils

import (
	"strings"
)

func CheckPBN(pbn string) (bool, error) {
	pbn = pbn[2:]

	parts := strings.Split(pbn, " ")
	if len(parts) != 4 {
		return false, INVALIDPARTS
	}

	var (
		Cards   = "AKQJT98765432"
		Spades  = make(map[string]struct{}, 13)
		Hearts  = make(map[string]struct{}, 13)
		Diamods = make(map[string]struct{}, 13)
		Clubs   = make(map[string]struct{}, 13)
	)

	for _, part := range parts {
		suits := strings.Split(part, ".")
		if len(suits) != 4 {
			return false, INVALIDSUITS
		}

		for i, suit := range suits {
			for _, card := range suit {
				if !strings.Contains(Cards, string(card)) {
					return false, INVALIDCARD
				}

				switch i {
				case 0:
					Spades[string(card)] = struct{}{}

				case 1:
					Hearts[string(card)] = struct{}{}

				case 2:
					Diamods[string(card)] = struct{}{}

				case 3:
					Clubs[string(card)] = struct{}{}
				}
			}
		}
	}

	if len(Spades) != 13 || len(Hearts) != 13 ||
		len(Diamods) != 13 || len(Clubs) != 13 {
		return false, INVALIDCOUNTOFCARDS
	}

	return true, nil
}
