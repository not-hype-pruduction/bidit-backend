package systems

import (
	"context"

	"github.com/not-hype-pruduction/bridge-backend/internal/domain/biding"
	"github.com/not-hype-pruduction/bridge-backend/internal/lib/utils"
)

type Language string

const (
	RU Language = "ru"
	EN Language = "en"
)

type SputnikStandard struct {
	Lang Language
}

func (s *SputnikStandard) Name() string {
	return "SYSTEM_SPUTNIK"
}

// вспомогательный метод для выбора текста
func (s *SputnikStandard) msg(ru, en string) string {
	if s.Lang == EN {
		return en
	}
	return ru
}

type handAnalysis struct {
	hcp    int
	counts map[biding.Suit]int
}

func (s *SputnikStandard) GetBid(ctx context.Context, in biding.BidState) (biding.Call, error) {
	analysis := s.analyze(utils.PBNToSlice(in.Hand))
	history := in.AuctionHistory

	if s.isOpening(history) {
		return s.open(analysis), nil
	}

	if len(history) >= 1 && s.partnerOpened(history) {
		partnerOpening := s.getLastNonPassBid(history)
		return s.respond(analysis, partnerOpening, history), nil
	}

	return biding.Call{Type: "PASS", Explanation: s.msg("Пас", "Pass")}, nil
}

// --- ЛОГИКА ОТКРЫТИЙ ---

func (s *SputnikStandard) open(info handAnalysis) biding.Call {
	hcp := info.hcp

	// 2БК: 20-21
	if hcp >= 20 && hcp <= 21 && s.isBalanced(info) {
		return biding.Call{
			Level: 2, Suit: biding.NoTrump, Type: "BID",
			Explanation: s.msg("20-21 очков, равномерный расклад", "20-21 HCP, balanced hand"),
		}
	}

	// 2К: 22+
	if hcp >= 22 {
		return biding.Call{
			Level: 2, Suit: biding.Clubs, Type: "BID",
			Explanation: s.msg("Сильное открытие 22+ очков, форсинг гейм", "Strong opening 22+ HCP, Game Forcing"),
		}
	}

	// 1БК: 15-17
	if hcp >= 15 && hcp <= 17 && s.isBalanced(info) && info.counts[biding.Spades] < 5 && info.counts[biding.Hearts] < 5 {
		return biding.Call{
			Level: 1, Suit: biding.NoTrump, Type: "BID",
			Explanation: s.msg("15-17 очков, равномерный расклад", "15-17 HCP, balanced hand"),
		}
	}

	// Блоки 3-4 уровня
	if hcp >= 5 && hcp <= 9 {
		for _, suit := range []biding.Suit{biding.Spades, biding.Hearts, biding.Diamonds, biding.Clubs} {
			if info.counts[suit] >= 7 {
				level := 3
				if info.counts[suit] >= 8 {
					level = 4
				}
				return biding.Call{
					Level: level, Suit: suit, Type: "BID",
					Explanation: s.msg("Блок: длинная масть, 5-9 очков", "Preemptive: long suit, 5-9 HCP"),
				}
			}
		}
	}

	// Слабые 2
	if hcp >= 5 && hcp <= 10 {
		for _, suit := range []biding.Suit{biding.Spades, biding.Hearts, biding.Diamonds} {
			if info.counts[suit] == 6 {
				return biding.Call{
					Level: 2, Suit: suit, Type: "BID",
					Explanation: s.msg("Слабое открытие: 6 карт, 5-10 очков", "Weak 2 opening: 6 cards, 5-10 HCP"),
				}
			}
		}
	}

	// Открытия 1 масть (12-21)
	if hcp >= 12 {
		if info.counts[biding.Spades] >= 5 {
			return biding.Call{Level: 1, Suit: biding.Spades, Type: "BID", Explanation: s.msg("5+ пик, 12-21 очков", "5+ Spades, 12-21 HCP")}
		}
		if info.counts[biding.Hearts] >= 5 {
			return biding.Call{Level: 1, Suit: biding.Hearts, Type: "BID", Explanation: s.msg("5+ черв, 12-21 очков", "5+ Hearts, 12-21 HCP")}
		}

		dCount := info.counts[biding.Diamonds]
		cCount := info.counts[biding.Clubs]
		if dCount >= 3 && (dCount > cCount || (dCount == 3 && cCount == 2)) {
			return biding.Call{Level: 1, Suit: biding.Diamonds, Type: "BID", Explanation: s.msg("3+ бубен, натуральное открытие", "3+ Diamonds, natural opening")}
		}
		return biding.Call{Level: 1, Suit: biding.Clubs, Type: "BID", Explanation: s.msg("3+ треф, натуральное открытие", "3+ Clubs, natural opening")}
	}

	return biding.Call{Type: "PASS"}
}

// --- ЛОГИКА ОТВЕТОВ ---

func (s *SputnikStandard) respond(info handAnalysis, opening biding.Call, history []biding.Call) biding.Call {
	if opening.Level == 1 && opening.Suit == biding.NoTrump {
		return s.respondTo1NT(info)
	}
	if opening.Level == 1 && (opening.Suit == biding.Hearts || opening.Suit == biding.Spades) {
		return s.respondToMajor(info, opening)
	}
	if opening.Level == 1 && (opening.Suit == biding.Clubs || opening.Suit == biding.Diamonds) {
		return s.respondToMinor(info, opening)
	}
	return biding.Call{Type: "PASS"}
}

func (s *SputnikStandard) respondToMinor(info handAnalysis, opening biding.Call) biding.Call {
	if info.hcp < 6 {
		return biding.Call{Type: "PASS"}
	}

	// 1 в мажоре
	if info.counts[biding.Hearts] >= 4 && (info.counts[biding.Hearts] >= info.counts[biding.Spades] || info.counts[biding.Spades] < 4) {
		return biding.Call{Level: 1, Suit: biding.Hearts, Type: "BID", Explanation: s.msg("6+ очков, 4+ черви, Ф1", "6+ HCP, 4+ Hearts, F1")}
	}
	if info.counts[biding.Spades] >= 4 {
		return biding.Call{Level: 1, Suit: biding.Spades, Type: "BID", Explanation: s.msg("6+ очков, 4+ пики, Ф1", "6+ HCP, 4+ Spades, F1")}
	}

	// БК ответы
	if info.hcp >= 6 && info.hcp <= 9 && s.isBalanced(info) {
		return biding.Call{Level: 1, Suit: biding.NoTrump, Type: "BID", Explanation: s.msg("6-9 очков, равномерный расклад", "6-9 HCP, balanced hand")}
	}
	if info.hcp >= 10 && info.hcp <= 11 && s.isBalanced(info) {
		return biding.Call{Level: 2, Suit: biding.NoTrump, Type: "BID", Explanation: s.msg("10-11 очков, инвит к 3БК", "10-11 HCP, invite to 3NT")}
	}

	return biding.Call{Type: "PASS"}
}

func (s *SputnikStandard) respondToMajor(info handAnalysis, opening biding.Call) biding.Call {
	if info.hcp < 6 {
		return biding.Call{Type: "PASS"}
	}

	// Фит
	if info.counts[opening.Suit] >= 3 {
		if info.hcp >= 6 && info.hcp <= 9 {
			return biding.Call{Level: 2, Suit: opening.Suit, Type: "BID", Explanation: s.msg("6-9 очков, поддержка (фит)", "6-9 HCP, support (fit)")}
		}
		if info.hcp >= 10 && info.hcp <= 11 {
			return biding.Call{Level: 3, Suit: opening.Suit, Type: "BID", Explanation: s.msg("10-11 очков, инвит к гейму", "10-11 HCP, game invite")}
		}
	}

	// 2-в-1 ФГ
	if info.hcp >= 12 {
		return biding.Call{Level: 2, Suit: biding.Clubs, Type: "BID", Explanation: s.msg("12+ очков, форсинг гейм", "12+ HCP, Game Forcing")}
	}

	return biding.Call{Level: 1, Suit: biding.NoTrump, Type: "BID", Explanation: s.msg("6-9 очков, нет фита", "6-9 HCP, no fit")}
}

func (s *SputnikStandard) respondTo1NT(info handAnalysis) biding.Call {
	// Стейман
	if info.hcp >= 8 && (info.counts[biding.Hearts] == 4 || info.counts[biding.Spades] == 4) {
		return biding.Call{Level: 2, Suit: biding.Clubs, Type: "BID", Explanation: s.msg("Стейман: запрос мажорных четверок", "Stayman: asking for 4-card majors")}
	}

	// Трансферы
	if info.counts[biding.Hearts] >= 5 {
		return biding.Call{Level: 2, Suit: biding.Diamonds, Type: "BID", Explanation: s.msg("Трансфер в черву", "Transfer to Hearts")}
	}
	if info.counts[biding.Spades] >= 5 {
		return biding.Call{Level: 2, Suit: biding.Hearts, Type: "BID", Explanation: s.msg("Трансфер в пику", "Transfer to Spades")}
	}

	if info.hcp >= 10 {
		return biding.Call{Level: 3, Suit: biding.NoTrump, Type: "BID", Explanation: s.msg("10+ очков, постановка гейма", "10+ HCP, bidding game")}
	}

	return biding.Call{Type: "PASS"}
}

// --- ВСПОМОГАТЕЛЬНЫЕ МЕТОДЫ ---

func (s *SputnikStandard) analyze(cards []string) handAnalysis {
	info := handAnalysis{counts: make(map[biding.Suit]int)}
	weights := map[string]int{"A": 4, "K": 3, "Q": 2, "J": 1}

	for _, card := range cards {
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

func (s *SputnikStandard) isBalanced(info handAnalysis) bool {
	shortSuits := 0
	for _, count := range info.counts {
		if count <= 1 {
			return false
		}
		if count == 2 {
			shortSuits++
		}
	}
	return shortSuits <= 1
}

func (s *SputnikStandard) isOpening(history []biding.Call) bool {
	for _, c := range history {
		if c.Type != "PASS" {
			return false
		}
	}
	return true
}

func (s *SputnikStandard) partnerOpened(history []biding.Call) bool {
	count := 0
	var last biding.Call
	for _, c := range history {
		if c.Type == "BID" {
			count++
			last = c
		}
	}
	return count == 1 && last.Type == "BID"
}

func (s *SputnikStandard) getLastNonPassBid(history []biding.Call) biding.Call {
	for i := len(history) - 1; i >= 0; i-- {
		if history[i].Type == "BID" {
			return history[i]
		}
	}
	return biding.Call{Type: "PASS"}
}
