package core

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
)

// Contact contact core object definition
type Contact struct {
	Id           string `db:"id"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	PhoneNumbers []Number
	Emails       []Email
}

func (c Contact) ToContactResponseDto() *dto.ContactResponse {
	response := new(dto.ContactResponse)

	response.Id = c.Id
	response.FirstName = c.FirstName
	response.LastName = c.LastName

	for _, number := range c.PhoneNumbers {
		response.PhoneNumbers = append(response.PhoneNumbers, dto.PhoneNumberResp{
			Id:        number.Id,
			ContactId: number.ContactId,
			Number:    number.PhoneNumber,
			Label:     number.Label,
		})
	}

	for _, email := range c.Emails {
		response.Emails = append(response.Emails, dto.EmailResp{
			Id:        email.Id,
			ContactId: email.ContactId,
			Address:   email.Address,
		})
	}

	return response
}

// ContactRepository Contact secondary port
type ContactRepository interface {
	CreateContact(*Contact) (*Contact, *errs.AppError)
}
