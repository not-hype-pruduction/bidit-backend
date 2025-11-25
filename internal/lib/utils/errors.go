package utils

import "errors"

var (
	INVALIDPARTS        = errors.New("the count of parts in pbn is invalid")
	INVALIDSUITS        = errors.New("the invalid count of suits in hand")
	INVALIDCARD         = errors.New("invalid card in suit")
	INVALIDCOUNTOFCARDS = errors.New("invalid count of cards in a suit")
)
