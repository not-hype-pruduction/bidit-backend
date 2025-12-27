package biding

type Suit int

const (
	Clubs Suit = iota + 1
	Diamonds
	Hearts
	Spades
	NoTrump
)

type Call struct {
	Level int
	Suit  Suit
	Type  string // PASS, DOUBLE, REDOUBLE, BID
}

type BidState struct {
	Hand           string
	AuctionHistory []Call
	Vulnerability  string
	SystemName     string
}
