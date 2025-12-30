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

// handAnalysis содержит посчитанные параметры руки
type handAnalysis struct {
	hcp    int
	counts map[biding.Suit]int
}

func (s *SputnikStandard) GetBid(ctx context.Context, in biding.BidState) (biding.Call, error) {
	analysis := s.analyze(utils.PBNToSlice(in.Hand))
	history := in.AuctionHistory

	// 1. Если мы открываем (никто еще не заявлялся кроме Паса)
	if s.isOpening(history) {
		return s.open(analysis), nil
	}

	// 2. Если партнер открылся, а оппоненты пасуют (Ответы)
	if len(history) >= 1 && s.partnerOpened(history) {
		partnerOpening := s.getLastNonPassBid(history)
		return s.respond(analysis, partnerOpening, history), nil
	}

	// По умолчанию ПАС
	return biding.Call{Type: "PASS"}, nil
}

// --- ЛОГИКА ОТКРЫТИЙ (Таблица 1) ---

func (s *SputnikStandard) open(info handAnalysis) biding.Call {
	hcp := info.hcp

	// 2БК: 20-21, равномерная
	if hcp >= 20 && hcp <= 21 && s.isBalanced(info) {
		return biding.Call{Level: 2, Suit: biding.NoTrump, Type: "BID"}
	}

	// 2К: 22+ очков, любая рука (ФГ)
	if hcp >= 22 {
		return biding.Call{Level: 2, Suit: biding.Clubs, Type: "BID"}
	}

	// 1БК: 15-17, равномерная, нет мажора 5+
	if hcp >= 15 && hcp <= 17 && s.isBalanced(info) && info.counts[biding.Spades] < 5 && info.counts[biding.Hearts] < 5 {
		return biding.Call{Level: 1, Suit: biding.NoTrump, Type: "BID"}
	}

	// Блоки на 3 и 4 уровне (5-9 очков)
	if hcp >= 5 && hcp <= 9 {
		for _, suit := range []biding.Suit{biding.Spades, biding.Hearts, biding.Diamonds, biding.Clubs} {
			if info.counts[suit] == 8 {
				return biding.Call{Level: 4, Suit: suit, Type: "BID"}
			}
			if info.counts[suit] == 7 {
				return biding.Call{Level: 3, Suit: suit, Type: "BID"}
			}
		}
	}

	// Слабые 2 (5-10 очков, 6 карт)
	if hcp >= 5 && hcp <= 10 {
		for _, suit := range []biding.Suit{biding.Spades, biding.Hearts, biding.Diamonds} {
			if info.counts[suit] == 6 {
				return biding.Call{Level: 2, Suit: suit, Type: "BID"}
			}
		}
	}

	// Открытия от 12 очков
	if hcp >= 12 {
		// Мажоры от 5 карт (Пика приоритетнее при 5-5)
		if info.counts[biding.Spades] >= 5 {
			return biding.Call{Level: 1, Suit: biding.Spades, Type: "BID"}
		}
		if info.counts[biding.Hearts] >= 5 {
			return biding.Call{Level: 1, Suit: biding.Hearts, Type: "BID"}
		}

		// Миноры (3+ карты)
		// Правило: 1Бубна если 4-4 в минорах (больше очков в бубне) или 3бубны-2трефы
		dCount := info.counts[biding.Diamonds]
		cCount := info.counts[biding.Clubs]

		if dCount > cCount {
			return biding.Call{Level: 1, Suit: biding.Diamonds, Type: "BID"}
		}
		if dCount == cCount && dCount >= 4 {
			// В системе Спутник при 4-4 в минорах выбор зависит от качества,
			// упростим: 1Т при 4-4, если не указано иное
			return biding.Call{Level: 1, Suit: biding.Clubs, Type: "BID"}
		}
		if dCount == 3 && cCount == 3 {
			return biding.Call{Level: 1, Suit: biding.Clubs, Type: "BID"}
		}
		if dCount == 3 && cCount == 2 {
			return biding.Call{Level: 1, Suit: biding.Diamonds, Type: "BID"}
		}

		return biding.Call{Level: 1, Suit: biding.Clubs, Type: "BID"}
	}

	return biding.Call{Type: "PASS"}
}

// --- ЛОГИКА ОТВЕТОВ (Таблицы 2, 3, 4) ---

func (s *SputnikStandard) respond(info handAnalysis, opening biding.Call, history []biding.Call) biding.Call {
	// Если партнер открылся 1БК
	if opening.Level == 1 && opening.Suit == biding.NoTrump {
		return s.respondTo1NT(info)
	}

	// Если партнер открылся 1 в мажоре
	if opening.Level == 1 && (opening.Suit == biding.Hearts || opening.Suit == biding.Spades) {
		return s.respondToMajor(info, opening)
	}

	// Если партнер открылся 1 в миноре
	if opening.Level == 1 && (opening.Suit == biding.Clubs || opening.Suit == biding.Diamonds) {
		return s.respondToMinor(info, opening)
	}

	return biding.Call{Type: "PASS"}
}

func (s *SputnikStandard) respondToMinor(info handAnalysis, opening biding.Call) biding.Call {
	if info.hcp < 6 {
		return biding.Call{Type: "PASS"}
	}

	// Ответы 1 в мажоре (4+ карты, приоритет Черве при 4-4)
	if info.counts[biding.Hearts] >= 4 && (info.counts[biding.Hearts] >= info.counts[biding.Spades] || info.counts[biding.Spades] < 4) {
		return biding.Call{Level: 1, Suit: biding.Hearts, Type: "BID"}
	}
	if info.counts[biding.Spades] >= 4 {
		return biding.Call{Level: 1, Suit: biding.Spades, Type: "BID"}
	}

	// БК ответы
	if info.hcp >= 6 && info.hcp <= 9 && s.isBalanced(info) {
		return biding.Call{Level: 1, Suit: biding.NoTrump, Type: "BID"}
	}
	if info.hcp >= 10 && info.hcp <= 11 && s.isBalanced(info) {
		return biding.Call{Level: 2, Suit: biding.NoTrump, Type: "BID"}
	}
	if info.hcp >= 12 && info.hcp <= 15 && s.isBalanced(info) {
		return biding.Call{Level: 3, Suit: biding.NoTrump, Type: "BID"}
	}

	return biding.Call{Type: "PASS"}
}

func (s *SputnikStandard) respondToMajor(info handAnalysis, opening biding.Call) biding.Call {
	if info.hcp < 6 {
		return biding.Call{Type: "PASS"}
	}

	// Фит в мажоре партнера (3+ карты)
	if info.counts[opening.Suit] >= 3 {
		if info.hcp >= 6 && info.hcp <= 9 {
			return biding.Call{Level: 2, Suit: opening.Suit, Type: "BID"}
		}
		// Инвит 10-11 (через 2БК или форсирующий БК в зависимости от вариации,
		// по PDF: 3 в мажор = инвит)
		if info.hcp >= 10 && info.hcp <= 11 {
			return biding.Call{Level: 3, Suit: opening.Suit, Type: "BID"}
		}
	}

	// 1 Пика на 1 Черву
	if opening.Suit == biding.Hearts && info.counts[biding.Spades] >= 4 {
		return biding.Call{Level: 1, Suit: biding.Spades, Type: "BID"}
	}

	// 2-в-1 ФГ (12+ очков, новая масть на 2 уровне)
	if info.hcp >= 12 {
		if info.counts[biding.Clubs] >= 4 {
			return biding.Call{Level: 2, Suit: biding.Clubs, Type: "BID"}
		}
		if info.counts[biding.Diamonds] >= 4 {
			return biding.Call{Level: 2, Suit: biding.Diamonds, Type: "BID"}
		}
	}

	// 1БК (6-9 очков, не форсирующий по системе Спутник-Стандарт)
	return biding.Call{Level: 1, Suit: biding.NoTrump, Type: "BID"}
}

func (s *SputnikStandard) respondTo1NT(info handAnalysis) biding.Call {
	// Стейман (8+ очков, есть 4-ка в мажоре)
	if info.hcp >= 8 && (info.counts[biding.Hearts] == 4 || info.counts[biding.Spades] == 4) {
		return biding.Call{Level: 2, Suit: biding.Clubs, Type: "BID"}
	}

	// Трансферы (от 0 очков, 5+ в мажоре)
	if info.counts[biding.Hearts] >= 5 {
		return biding.Call{Level: 2, Suit: biding.Diamonds, Type: "BID"} // Трансфер в Черву
	}
	if info.counts[biding.Spades] >= 5 {
		return biding.Call{Level: 2, Suit: biding.Hearts, Type: "BID"} // Трансфер в Пику
	}

	// Инвит в 3БК (8-9 очков)
	if info.hcp >= 8 && info.hcp <= 9 {
		return biding.Call{Level: 2, Suit: biding.NoTrump, Type: "BID"}
	}

	// Гейм 3БК (10-15 очков)
	if info.hcp >= 10 && info.hcp <= 15 {
		return biding.Call{Level: 3, Suit: biding.NoTrump, Type: "BID"}
	}

	return biding.Call{Type: "PASS"}
}

func (s *SputnikStandard) analyze(cards []string) handAnalysis {
	info := handAnalysis{counts: make(map[biding.Suit]int)}
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

func (s *SputnikStandard) isBalanced(info handAnalysis) bool {
	// Равномерная рука: нет синглетов/ренонсов, не более одного дублета
	shortSuits := 0
	for _, count := range info.counts {
		if count <= 1 {
			return false // Есть синглет или ренонс
		}
		if count == 2 {
			shortSuits++
		}
	}
	return shortSuits <= 1
}

func (s *SputnikStandard) isOpening(history []biding.Call) bool {
	if len(history) == 0 {
		return true
	}
	for _, c := range history {
		if c.Type != "PASS" {
			return false
		}
	}
	return true
}

func (s *SputnikStandard) partnerOpened(history []biding.Call) bool {
	// Ищем последнюю значащую заявку. Если она сделана партнером (через одного от нас)
	// Для простоты: если в истории всего одна значащая заявка и она не наша.
	count := 0
	var lastBid biding.Call
	for _, c := range history {
		if c.Type == "BID" {
			count++
			lastBid = c
		}
	}
	// Если мы в позиции отвечающего (ход 2 или 4)
	return count == 1 && lastBid.Type == "BID"
}

func (s *SputnikStandard) getLastNonPassBid(history []biding.Call) biding.Call {
	for i := len(history) - 1; i >= 0; i-- {
		if history[i].Type == "BID" {
			return history[i]
		}
	}
	return biding.Call{Type: "PASS"}
}
