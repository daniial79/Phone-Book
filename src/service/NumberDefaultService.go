package service

import (
	"github.com/daniial79/Phone-Book/src/core"
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
)

// NumberDefaultService primary actor
type NumberDefaultService struct {
	repo core.NumberRepository
}

func NewNumberDefaultService(repo core.NumberRepository) NumberDefaultService {
	return NumberDefaultService{repo}
}

func (s NumberDefaultService) AddNewNumbers(request dto.AddNumberRequest) (*dto.AddNumberResponse, *errs.AppError) {
	coreTypedNumber := core.Number{
		Id:          "",
		ContactId:   request.ContactId,
		PhoneNumber: request.Number,
		Label:       request.Label,
	}

	addedNumber, err := s.repo.AddNewNumber(coreTypedNumber)
	if err != nil {
		return nil, err
	}

	response := addedNumber.ToAddNumberResponseDto()

	return response, nil
}
