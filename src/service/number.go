package service

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
)

// NumberService primary port
type NumberService interface {
	AddNewNumbers(dto.AddNumberRequest) (*dto.AddNumberResponse, *errs.AppError)
}
