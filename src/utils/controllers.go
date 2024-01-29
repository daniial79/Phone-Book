package utils

import (
	"fmt"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/google/uuid"
)

func ConvertStringToUUID(bytes []byte) (uuid.UUID, *errs.AppError) {
	generatedUUID, err := uuid.FromBytes(bytes)
	fmt.Println(err)
	if err != nil {
		return uuid.UUID{}, errs.NewBadRequestErr(errs.InvalidIdErr)
	}

	return generatedUUID, nil
}
