package service

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
)

// ContactService Primary port
type ContactService interface {
	NewContact(username string, requestBody dto.NewContactRequest) (*dto.NewContactResponse, *errs.AppError)
	AddNewNumbers(requestBody []dto.AddNumberRequest, contactId string) ([]dto.AddNumberResponse, *errs.AppError)
	AddNewEmails(requestBody []dto.AddEmailRequest, contactId string) ([]dto.AddEmailResponse, *errs.AppError)
	GetContacts(username string) ([]dto.NewContactResponse, *errs.AppError)
	GetContactCredentials(cId string) (*dto.ContactCredentialsResponse, *errs.AppError)
	DeleteEmailFromContact(cId, eId string) (*dto.NoContentResponse, *errs.AppError)
	DeletePhoneNumberFromContact(cId, eId string) (*dto.NoContentResponse, *errs.AppError)
	DeleteContact(cId string) (*dto.NoContentResponse, *errs.AppError)
	UpdateContactNumber(cId, nId string, requestBody dto.UpdateNumberRequest) (*dto.UpdateNumberResponse, *errs.AppError)
	UpdateContactEmail(cId, eId string, requestBody dto.UpdateEmailRequest) (*dto.UpdateEmailResponse, *errs.AppError)
	UpdateContact(cId string, requestBody dto.UpdateContactRequest) (*dto.UpdateContactResponse, *errs.AppError)
}
