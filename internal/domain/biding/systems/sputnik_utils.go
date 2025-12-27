package systems

import "github.com/not-hype-pruduction/bridge-backend/internal/domain/biding"

// Аналитика руки
type handInfo struct {
	hcp    int
	counts map[biding.Suit]int
}

func analyzeHand(cards []string) handInfo {
	info := handInfo{counts: make(map[biding.Suit]int)}
	weights := map[string]int{"A": 4, "K": 3, "Q": 2, "J": 1}

	for _, card := range cards {
		// Формат "SA" (Spades Ace), "H10" (Hearts 10)
		suitChar := string(card[0])
		rank := card[1:]

		var suit biding.Suit
		switch suitChar {
		case "S":
			suit = biding.Spades
		case "H":
			suit = biding.Hearts
		case "D":
			suit = biding.Diamonds
		case "C":
			suit = biding.Clubs
		}

		info.counts[suit]++
		if val, ok := weights[rank]; ok {
			info.hcp += val
		}
	}
	return info
}

func (h handInfo) isBalanced() bool {
	// Равномерные руки: 4333, 4432, 5332
	// Упрощенно: нет синглетов и ренонсов, не более одной 2-картной масти
	twos := 0
	for _, count := range h.counts {
		if count < 2 {
			return false
		}
		if count == 2 {
			twos++
		}
	}
	return twos <= 1
}
