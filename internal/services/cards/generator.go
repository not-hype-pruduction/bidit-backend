package cards

import (
	"fmt"
	"math/rand"
)

func generateDeal(gCards generatedCards, dealer string) *deal {
	var resDeal deal

	resDeal.Dealer = dealer

	aces := []card{
		{Rank: "A", Suit: "S"},
		{Rank: "A", Suit: "H"},
		{Rank: "A", Suit: "D"},
		{Rank: "A", Suit: "C"},
	}

	kings := []card{
		{Rank: "K", Suit: "S"},
		{Rank: "K", Suit: "H"},
		{Rank: "K", Suit: "D"},
		{Rank: "K", Suit: "C"},
	}

	queens := []card{
		{Rank: "Q", Suit: "S"},
		{Rank: "Q", Suit: "H"},
		{Rank: "Q", Suit: "D"},
		{Rank: "Q", Suit: "C"},
	}

	jacks := []card{
		{Rank: "J", Suit: "S"},
		{Rank: "J", Suit: "H"},
		{Rank: "J", Suit: "D"},
		{Rank: "J", Suit: "C"},
	}

	var numberCards = []card{
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
	assignCards(&resDeal.My, numberCards, &numbersIndex, numbersIndex+13-resDeal.My.getAmountOfCards())
	assignCards(&resDeal.Partner, numberCards, &numbersIndex, numbersIndex+13-resDeal.Partner.getAmountOfCards())
	assignCards(&resDeal.Player1, numberCards, &numbersIndex, numbersIndex+13-resDeal.Player1.getAmountOfCards())
	assignCards(&resDeal.Player2, numberCards, &numbersIndex, numbersIndex+13-resDeal.Player2.getAmountOfCards())

	fmt.Println(resDeal)

	return &resDeal
}

func assignCards(gHand *hand, cards []card, index *int, count int) {
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
