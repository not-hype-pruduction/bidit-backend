// TODO: need to imlement all system
package systems

import (
	"context"

	"github.com/not-hype-pruduction/bridge-backend/internal/domain/biding"
	"github.com/not-hype-pruduction/bridge-backend/internal/lib/utils"
)

type SputnikStandard struct{}

func (s *SputnikStandard) Name() string {
	return "SYSTEM_SPUTNIK"
}

func (s *SputnikStandard) GetBid(ctx context.Context, in biding.BidState) (biding.Call, error) {
	info := analyzeHand(utils.PBNToSlice(in.Hand))
	history := in.AuctionHistory

	if len(history) == 0 || isAllPass(history) {
		return s.open(info)
	}

	if len(history) >= 1 {
		lastBid := history[len(history)-1]
		// Логика ответов на 1 в миноре (стр. 1 PDF)
		if lastBid.Level == 1 && (lastBid.Suit == biding.Clubs || lastBid.Suit == biding.Diamonds) {
			return s.respondToMinor(info, lastBid)
		}
	}

	return biding.Call{Type: "PASS"}, nil
}

// Логика открытий (Таблица 1 из PDF)
func (s *SputnikStandard) open(info handInfo) (biding.Call, error) {
	hcp := info.hcp

	// Блоки
	if hcp >= 5 && hcp <= 10 {
		if info.counts[biding.Spades] == 6 {
			return biding.Call{Level: 2, Suit: biding.Spades, Type: "BID"}, nil
		}
		if info.counts[biding.Hearts] == 6 {
			return biding.Call{Level: 2, Suit: biding.Hearts, Type: "BID"}, nil
		}
		if info.counts[biding.Diamonds] == 6 {
			return biding.Call{Level: 2, Suit: biding.Diamonds, Type: "BID"}, nil
		}
	}

	// Открытия от 12 очков
	if hcp < 12 {
		return biding.Call{Type: "PASS"}, nil
	}

	// 1БК (15-17, равномерная)
	if hcp >= 15 && hcp <= 17 && info.isBalanced() {
		return biding.Call{Level: 1, Suit: biding.NoTrump, Type: "BID"}, nil
	}

	// Мажоры от 5 карт
	if info.counts[biding.Spades] >= 5 {
		return biding.Call{Level: 1, Suit: biding.Spades, Type: "BID"}, nil
	}
	if info.counts[biding.Hearts] >= 5 {
		return biding.Call{Level: 1, Suit: biding.Hearts, Type: "BID"}, nil
	}

	// Миноры (правило выбора между 1C и 1D)
	if info.counts[biding.Diamonds] >= 4 || (info.counts[biding.Diamonds] == 3 && info.counts[biding.Clubs] == 2) {
		return biding.Call{Level: 1, Suit: biding.Diamonds, Type: "BID"}, nil
	}

	return biding.Call{Level: 1, Suit: biding.Clubs, Type: "BID"}, nil
}

// Логика ответов на 1 минор (Таблица 2 из PDF)
func (s *SputnikStandard) respondToMinor(info handInfo, opening biding.Call) (biding.Call, error) {
	if info.hcp < 6 {
		return biding.Call{Type: "PASS"}, nil
	}

	// Приоритет мажорам от 4 карт
	if info.counts[biding.Hearts] >= 4 && info.counts[biding.Hearts] >= info.counts[biding.Spades] {
		return biding.Call{Level: 1, Suit: biding.Hearts, Type: "BID"}, nil
	}
	if info.counts[biding.Spades] >= 4 {
		return biding.Call{Level: 1, Suit: biding.Spades, Type: "BID"}, nil
	}

	// 1БК ответ (6-9 очков, без мажоров)
	if info.hcp <= 9 {
		return biding.Call{Level: 1, Suit: biding.NoTrump, Type: "BID"}, nil
	}

	return biding.Call{Type: "PASS"}, nil
}

func isAllPass(history []biding.Call) bool {
	for _, call := range history {
		if call.Type != "PASS" {
			return false
		}
	}
	return true
}
