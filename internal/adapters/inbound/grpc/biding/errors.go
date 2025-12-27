// Package biding contains the gRPC handler for biding part of game.
package biding

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidBidLevel = status.Error(codes.InvalidArgument, "invalid bid level")
)
