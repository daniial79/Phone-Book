package service

import (
	"github.com/daniial79/Phone-Book/src/core"
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
)

// EmailDefaultService Primary Actor
type EmailDefaultService struct {
	repo core.EmailRepository
}

func NewEmailDefaultService(repo core.EmailRepository) EmailDefaultService {
	return EmailDefaultService{repo}
}

func (s EmailDefaultService) AddNewEmails(request []dto.AddEmailRequest) ([]dto.AddEmailResponse, *errs.AppError) {
	coreTypedEmails := make([]core.Email, len(request))

	for i, email := range request {
		coreTypedEmails[i] = core.Email{
			Id:        "",
			ContactId: email.ContactId,
			Address:   email.Address,
		}
	}

	addedEmails, err := s.repo.AddNewEmails(coreTypedEmails)
	if err != nil {
		return nil, err
	}

	response := make([]dto.AddEmailResponse, len(coreTypedEmails))
	for i, email := range addedEmails {
		response[i] = email.ToAddEmailResponseDto()
	}

	return response, nil
}
