package cards

import (
	"context"
	"errors"
	"log/slog"

	cardsv1 "github.com/not-hype-pruduction/bridge-backend/internal/pb/cards.v1"
)

var (
	IMPOSSIBLECARDCOMBINATION = errors.New("The selected card combination cannot occur.")
)

type Cards struct {
	log    *slog.Logger
	random bool
}

func New(log *slog.Logger, random bool) Cards {
	return Cards{
		log:    log,
		random: random,
	}
}

func (c *Cards) GenerateHands(
	ctx context.Context,
	myPointsMin int32,
	myPointsMax int32,
	partnerPointsMin int32,
	partnerPointsMax int32,
	dealer string,
	north cardsv1.User,
	south cardsv1.User,
	west cardsv1.User,
	east cardsv1.User,
) (string, error) {
	cards, err := generateCardsWithPoints(
		myPointsMin, myPointsMax,
		partnerPointsMin, partnerPointsMax,
	)
	if err != nil {
		return "", err
	}

	c.log.Debug("", *cards)

	deal := generateDeal(*cards, dealer)

	return deal.createPBN(int32(north), int32(south), int32(east), int32(west))
}
