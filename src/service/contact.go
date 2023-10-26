package service

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
)

// ContactService Primary port
type ContactService interface {
	NewContact(request dto.NewContactRequest) (*dto.NewContactResponse, *errs.AppError)
	AddNewNumbers(request []dto.AddNumberRequest, contactId string) ([]dto.AddNumberResponse, *errs.AppError)
	AddNewEmails(request []dto.AddEmailRequest, contactId string) ([]dto.AddEmailResponse, *errs.AppError)
	GetContacts() ([]dto.NewContactResponse, *errs.AppError)
	GetContactCredentials(cId string) (*dto.ContactCredentialsResponse, *errs.AppError)
	DeleteEmailFromContact(cId, eId string) (*dto.NoContentResponse, *errs.AppError)
	DeletePhoneNumberFromContact(cId, eId string) (*dto.NoContentResponse, *errs.AppError)
}
