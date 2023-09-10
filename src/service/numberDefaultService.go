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

func (s NumberDefaultService) AddNewNumbers(request []dto.AddNumberRequest) ([]dto.AddNumberResponse, *errs.AppError) {
	coreTypedNumbers := make([]core.Number, len(request))

	for i, numberRequest := range request {
		coreTypedNumbers[i] = core.Number{
			Id:          "",
			ContactId:   numberRequest.ContactId,
			PhoneNumber: numberRequest.Number,
			Label:       numberRequest.Label,
		}
	}

	addedNumbers, err := s.repo.AddNewNumber(coreTypedNumbers)
	if err != nil {
		return nil, err
	}

	response := make([]dto.AddNumberResponse, len(addedNumbers))
	for i, number := range addedNumbers {
		response[i] = number.ToAddNumberResponseDto()
	}

	return response, nil
}
