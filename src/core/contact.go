package core

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/google/uuid"
)

// ContactRepository Contact secondary port
type ContactRepository interface {
	GetContactOwnerByUsername(username string) (uuid.UUID, *errs.AppError)
	CreateContact(string, *Contact) (*Contact, *errs.AppError)
	CheckContactExistenceById(cId uuid.UUID) *errs.AppError
	AddNewNumber(n []Number) ([]Number, *errs.AppError)
	AddNewEmails(e []Email) ([]Email, *errs.AppError)
	GetAllContacts(username string) ([]Contact, *errs.AppError)
	GetContactNumbers(cId uuid.UUID) ([]Number, *errs.AppError)
	GetContactEmails(cId uuid.UUID) ([]Email, *errs.AppError)
	DeleteContactEmail(cId, eId uuid.UUID) *errs.AppError
	DeleteContactPhoneNumber(cId, nId uuid.UUID) *errs.AppError
	DeleteContact(cId uuid.UUID) *errs.AppError
	UpdateContactPhoneNumber(newNumber Number) (*Number, *errs.AppError)
	UpdateContactEmail(newEmail Email) (*Email, *errs.AppError)
	UpdateContact(newContact Contact) (*Contact, *errs.AppError)
}

// Contact contact core object definition
type Contact struct {
	Id           uuid.UUID `db:"id"`
	OwnerId      uuid.UUID `db:"owner_id"`
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
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

func (c Contact) ToUpdatedContactResponseDto() dto.UpdateContactResponse {
	return dto.UpdateContactResponse{
		UpdatedFirstName: c.FirstName,
		UpdatedLastName:  c.LastName,
	}
}
