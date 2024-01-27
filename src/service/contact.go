package service

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/google/uuid"
)

// ContactService Primary port
type ContactService interface {
	NewContact(username string, requestBody dto.NewContactRequest) (*dto.NewContactResponse, *errs.AppError)
	AddNewNumbers(requestBody []dto.AddNumberRequest, contactId uuid.UUID) ([]dto.AddNumberResponse, *errs.AppError)
	AddNewEmails(requestBody []dto.AddEmailRequest, contactId uuid.UUID) ([]dto.AddEmailResponse, *errs.AppError)
	GetContacts(username string) ([]dto.NewContactResponse, *errs.AppError)
	GetContactCredentials(cId uuid.UUID) (*dto.ContactCredentialsResponse, *errs.AppError)
	DeleteEmailFromContact(cId, eId uuid.UUID) (*dto.NoContentResponse, *errs.AppError)
	DeletePhoneNumberFromContact(cId, eId uuid.UUID) (*dto.NoContentResponse, *errs.AppError)
	DeleteContact(cId uuid.UUID) (*dto.NoContentResponse, *errs.AppError)
	UpdateContactNumber(cId, nId uuid.UUID, requestBody dto.UpdateNumberRequest) (*dto.UpdateNumberResponse, *errs.AppError)
	UpdateContactEmail(cId, eId uuid.UUID, requestBody dto.UpdateEmailRequest) (*dto.UpdateEmailResponse, *errs.AppError)
	UpdateContact(cId uuid.UUID, requestBody dto.UpdateContactRequest) (*dto.UpdateContactResponse, *errs.AppError)
}
