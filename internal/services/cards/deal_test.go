package cards

import "testing"

func Test_cardsToPBNFormat(t *testing.T) {
	tests := []struct {
		name  string
		cards []card
		want  string
	}{
		{
			name: "simple",
			cards: []card{
				{Rank: "A", Suit: "S"},
				{Rank: "T", Suit: "S"},
				{Rank: "6", Suit: "S"},
			},
			want: "AT6",
		},
		{
			name: "reverse",
			cards: []card{
				{Rank: "6", Suit: "S"},
				{Rank: "T", Suit: "S"},
				{Rank: "A", Suit: "S"},
			},
			want: "AT6",
		},
		{
			name: "full",
			cards: []card{
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
			got := cardsToPBNFormat(tt.cards)
			if got != tt.want {
				t.Errorf("cardsToPBNFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handToPBNFormat(t *testing.T) {
	tests := []struct {
		name string
		hand *hand
		want string
	}{
		{
			name: "correst position",
			hand: &hand{
				Spades: []card{
					{Rank: "A", Suit: "S"},
				},
				Hearts: []card{
					{Rank: "K", Suit: "H"},
				},
				Diamonds: []card{
					{Rank: "Q", Suit: "D"},
				},
				Clubs: []card{
					{Rank: "J", Suit: "C"},
				},
			},
			want: "A.K.Q.J",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := handToPBNFormat(tt.hand)

			if got != tt.want {
				t.Errorf("handToPBNFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
