package cards_grpc

import (
	"context"
	"errors"

	cardsv1 "github.com/not-hype-pruduction/bridge-backend/internal/pb/cards.v1"
	"github.com/not-hype-pruduction/bridge-backend/internal/services/cards"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	POINTSERROR = errors.New("amount of points is incorrect, please provide a correct amount of points.")
	DEALERERROR = errors.New("the dealer is incorect, please provide a correct dealer.")
)

type serverAPI struct {
	cardsv1.UnimplementedCardGeneratorServiceServer
	cards Cards
}

type Cards interface {
	GenerateHands(
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
	) (pbn string, err error)
}

func Register(gPRCServer *grpc.Server, cards *Cards) {
	cardsv1.RegisterCardGeneratorServiceServer(gPRCServer, &serverAPI{cards: *cards})
}

func (s *serverAPI) GenerateHands(
	ctx context.Context,
	in *cardsv1.GenerateHandsRequest,
) (*cardsv1.GenerateHandsResponse, error) {
	if in.MyPointsMin > in.MyPointsMax ||
		in.PartnerPointsMin > in.PartnerPointsMax ||
		in.MyPointsMax+in.MyPointsMax > 40 {
		return nil, status.Error(codes.InvalidArgument, POINTSERROR.Error())
	}

	if in.Delaer != "S" && in.Delaer != "W" && in.Delaer != "N" && in.Delaer != "E" {
		return nil, status.Error(codes.InvalidArgument, DEALERERROR.Error())
	}

	pbn, err := s.cards.GenerateHands(
		ctx, in.MyPointsMin, in.MyPointsMax,
		in.PartnerPointsMin, in.PartnerPointsMax,
		in.Delaer, in.North, in.East,
		in.West, in.South,
	)
	if err != nil {
		if errors.Is(err, cards.IMPOSSIBLECARDCOMBINATION) {
			return nil, status.Error(codes.InvalidArgument, cards.IMPOSSIBLECARDCOMBINATION.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &cardsv1.GenerateHandsResponse{
		Pbn: pbn,
	}, nil
}
