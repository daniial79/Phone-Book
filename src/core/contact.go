package core

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
)

// ContactRepository Contact secondary port
type ContactRepository interface {
	CreateContact(*Contact) (*Contact, *errs.AppError)
	CheckContactExistenceById(cId string) *errs.AppError
	AddNewNumber(n []Number) ([]Number, *errs.AppError)
	AddNewEmails(e []Email) ([]Email, *errs.AppError)
	GetAllContacts() ([]Contact, *errs.AppError)
	GetContactNumbers(cId string) ([]Number, *errs.AppError)
	GetContactEmails(cId string) ([]Email, *errs.AppError)
	DeleteContactEmail(cId, eId string) *errs.AppError
	DeleteContactPhoneNumber(cId, nId string) *errs.AppError
	DeleteContact(cId string) *errs.AppError
	UpdateContactPhoneNumber(newNumber Number) (*Number, *errs.AppError)
	UpdateContactEmail(newEmail Email) (*Email, *errs.AppError)
}

// Contact contact core object definition
type Contact struct {
	Id           string `db:"id"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	PhoneNumbers []Number
	Emails       []Email
}

func (c Contact) ToContactResponseDto() *dto.NewContactResponse {
	response := new(dto.NewContactResponse)

	response.Id = c.Id
	response.FirstName = c.FirstName
	response.LastName = c.LastName

	for _, number := range c.PhoneNumbers {
		response.PhoneNumbers = append(response.PhoneNumbers, dto.PhoneNumberResponse{
			Id:          number.Id,
			ContactId:   number.ContactId,
			PhoneNumber: number.PhoneNumber,
			Label:       number.Label,
		})
	}

	for _, email := range c.Emails {
		response.Emails = append(response.Emails, dto.EmailResponse{
			Id:        email.Id,
			ContactId: email.ContactId,
			Address:   email.Address,
		})
	}

	return response
}
