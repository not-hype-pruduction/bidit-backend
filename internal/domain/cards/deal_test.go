package cards

import "testing"

func Test_CardsToPBNFormat(t *testing.T) {
	tests := []struct {
		name  string
		cards []Card
		want  string
	}{
		{
			name: "simple",
			cards: []Card{
				{Rank: "A", Suit: "S"},
				{Rank: "T", Suit: "S"},
				{Rank: "6", Suit: "S"},
			},
			want: "AT6",
		},
		{
			name: "reverse",
			cards: []Card{
				{Rank: "6", Suit: "S"},
				{Rank: "T", Suit: "S"},
				{Rank: "A", Suit: "S"},
			},
			want: "AT6",
		},
		{
			name: "full",
			cards: []Card{
				{Rank: "2", Suit: "S"},
				{Rank: "3", Suit: "S"},
				{Rank: "4", Suit: "S"},
				{Rank: "5", Suit: "S"},
				{Rank: "6", Suit: "S"},
				{Rank: "7", Suit: "S"},
				{Rank: "8", Suit: "S"},
				{Rank: "9", Suit: "S"},
				{Rank: "T", Suit: "S"},
				{Rank: "J", Suit: "S"},
				{Rank: "Q", Suit: "S"},
				{Rank: "K", Suit: "S"},
				{Rank: "A", Suit: "S"},
			},
			want: "AKQJT98765432",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CardsToPBNFormat(tt.cards)
			if got != tt.want {
				t.Errorf("CardsToPBNFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HandToPBNFormat(t *testing.T) {
	tests := []struct {
		name string
		hand *Hand
		want string
	}{
		{
			name: "correct position",
			hand: &Hand{
				Spades: []Card{
					{Rank: "A", Suit: "S"},
				},
				Hearts: []Card{
					{Rank: "K", Suit: "H"},
				},
				Diamonds: []Card{
					{Rank: "Q", Suit: "D"},
				},
				Clubs: []Card{
					{Rank: "J", Suit: "C"},
				},
			},
			want: "A.K.Q.J",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HandToPBNFormat(tt.hand)

			if got != tt.want {
				t.Errorf("HandToPBNFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
