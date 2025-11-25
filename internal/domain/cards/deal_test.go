package cards

import (
	"testing"

	"github.com/not-hype-pruduction/bridge-backend/internal/lib/utils"
)

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

func TestDeal_CreatePBN(t *testing.T) {
	tests := []struct {
		name string

		gCards  GeneratedCards
		dealer  string
		north   int32
		south   int32
		east    int32
		west    int32
		want    string
		wantErr bool
	}{
		{
			name: "standard distribution with all suits",
			gCards: GeneratedCards{
				A:  1,
				K:  2,
				Q:  3,
				J:  3,
				A_: 1,
				K_: 2,
				Q_: 1,
				J_: 1,
			},
			dealer:  "N",
			north:   1,
			south:   2,
			east:    3,
			west:    4,
			want:    "",
			wantErr: false,
		},
		{
			name: "dealer south",
			gCards: GeneratedCards{
				A:  4,
				K:  4,
				Q:  4,
				J:  4,
				A_: 0,
				K_: 0,
				Q_: 0,
				J_: 0,
			},
			dealer:  "S",
			north:   1,
			south:   2,
			east:    3,
			west:    4,
			wantErr: false,
		},
		{
			name: "dealer east",
			gCards: GeneratedCards{
				A:  2,
				K:  2,
				Q:  2,
				J:  2,
				A_: 2,
				K_: 2,
				Q_: 2,
				J_: 2,
			},
			dealer:  "E",
			north:   1,
			south:   2,
			east:    3,
			west:    4,
			wantErr: false,
		},
		{
			name: "dealer west",
			gCards: GeneratedCards{
				A:  3,
				K:  1,
				Q:  0,
				J:  4,
				A_: 1,
				K_: 3,
				Q_: 4,
				J_: 0,
			},
			dealer:  "W",
			north:   1,
			south:   2,
			east:    3,
			west:    4,
			wantErr: false,
		},
		{
			name: "no aces",
			gCards: GeneratedCards{
				A:  0,
				K:  4,
				Q:  4,
				J:  4,
				A_: 0,
				K_: 0,
				Q_: 0,
				J_: 0,
			},
			dealer:  "N",
			north:   1,
			south:   2,
			east:    3,
			west:    4,
			wantErr: false,
		},
		{
			name: "all aces to one player scenario",
			gCards: GeneratedCards{
				A:  4,
				K:  0,
				Q:  0,
				J:  0,
				A_: 0,
				K_: 4,
				Q_: 4,
				J_: 4,
			},
			dealer:  "S",
			north:   1,
			south:   2,
			east:    3,
			west:    4,
			wantErr: false,
		},
		{
			name: "balanced distribution",
			gCards: GeneratedCards{
				A:  1,
				K:  1,
				Q:  1,
				J:  1,
				A_: 1,
				K_: 1,
				Q_: 1,
				J_: 1,
			},
			dealer:  "N",
			north:   1,
			south:   2,
			east:    3,
			west:    4,
			wantErr: false,
		},
		{
			name: "extreme distribution - many high cards",
			gCards: GeneratedCards{
				A:  4,
				K:  4,
				Q:  4,
				J:  4,
				A_: 0,
				K_: 0,
				Q_: 0,
				J_: 0,
			},
			dealer:  "E",
			north:   1,
			south:   2,
			east:    3,
			west:    4,
			wantErr: false,
		},
		{
			name: "minimal high cards",
			gCards: GeneratedCards{
				A:  1,
				K:  1,
				Q:  1,
				J:  1,
				A_: 0,
				K_: 0,
				Q_: 0,
				J_: 0,
			},
			dealer:  "W",
			north:   1,
			south:   2,
			east:    3,
			west:    4,
			wantErr: false,
		},
		{
			name: "only jacks and queens",
			gCards: GeneratedCards{
				A:  0,
				K:  0,
				Q:  4,
				J:  4,
				A_: 0,
				K_: 0,
				Q_: 0,
				J_: 0,
			},
			dealer:  "N",
			north:   1,
			south:   2,
			east:    3,
			west:    4,
			wantErr: false,
		},
		{
			name: "mixed distribution with underscores",
			gCards: GeneratedCards{
				A:  2,
				K:  1,
				Q:  1,
				J:  0,
				A_: 2,
				K_: 3,
				Q_: 3,
				J_: 4,
			},
			dealer:  "S",
			north:   1,
			south:   2,
			east:    3,
			west:    4,
			wantErr: false,
		},
		{
			name: "equal distribution",
			gCards: GeneratedCards{
				A:  2,
				K:  2,
				Q:  2,
				J:  2,
				A_: 2,
				K_: 2,
				Q_: 2,
				J_: 2,
			},
			dealer:  "E",
			north:   1,
			south:   2,
			east:    3,
			west:    4,
			wantErr: false,
		},
		{
			name: "high board numbers",
			gCards: GeneratedCards{
				A:  3,
				K:  3,
				Q:  1,
				J:  1,
				A_: 1,
				K_: 1,
				Q_: 3,
				J_: 3,
			},
			dealer:  "W",
			north:   1,
			south:   2,
			east:    3,
			west:    4,
			wantErr: false,
		},
		{
			name: "dealer north with zero underscores",
			gCards: GeneratedCards{
				A:  4,
				K:  4,
				Q:  2,
				J:  1,
				A_: 0,
				K_: 0,
				Q_: 2,
				J_: 3,
			},
			dealer:  "N",
			north:   1,
			south:   2,
			east:    3,
			west:    4,
			wantErr: false,
		},
		{
			name: "asymmetric distribution",
			gCards: GeneratedCards{
				A:  1,
				K:  4,
				Q:  0,
				J:  2,
				A_: 3,
				K_: 0,
				Q_: 4,
				J_: 2,
			},
			dealer:  "S",
			north:   1,
			south:   2,
			east:    3,
			west:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := GenerateDeal(tt.gCards, tt.dealer)
			got, gotErr := d.CreatePBN(tt.north, tt.south, tt.east, tt.west)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreatePBN() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreatePBN() succeeded unexpectedly")
			}
			if ok, err := utils.CheckPBN(got); !ok {
				t.Errorf("CreatePBN() = %v, want %v", got, err)
			}
		})
	}
}
