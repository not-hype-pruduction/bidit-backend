package dds

import (
	"errors"

	"github.com/not-hype-pruduction/bridge-backend/internal/lib/utils"
	ddsv1 "github.com/not-hype-pruduction/bridge-backend/internal/pb/dds.v1"
)

func Validate(in *ddsv1.GetDDTableRequest) error {
	if err := validatePBN(in); err != nil {
		return errors.Join(ErrInvalidPBN, err)
	}

	return nil
}

func validatePBN(in *ddsv1.GetDDTableRequest) error {
	ok, err := utils.CheckPBN(in.Pbn)
	if !ok {
		return err
	}

	return nil
}
