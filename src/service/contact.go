package service

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
)

// ContactService Primary port
type ContactService interface {
	NewContact(request dto.NewContactRequest) (*dto.ContactResponse, *errs.AppError)
	AddNewNumbers(request []dto.AddNumberRequest, contactId string) ([]dto.AddNumberResponse, *errs.AppError)
	AddNewEmails(request []dto.AddEmailRequest, contactId string) ([]dto.AddEmailResponse, *errs.AppError)
	GetContacts(options map[string]string) ([]dto.ContactResponse, *errs.AppError)
}
