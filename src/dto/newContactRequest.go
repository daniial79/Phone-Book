package dto

import (
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/utils"
)

// PhoneNumberRequest this phone number request is for initiating new contact
type PhoneNumberRequest struct {
	Number string `json:"number"`
	Label  string `json:"label"`
}

// EmailRequest this email request is for initiating new contact
type EmailRequest struct {
	Address string `json:"address"`
}

// NewContactRequest dto object definition
type NewContactRequest struct {
	FirstName    string               `json:"firstName"`
	LastName     string               `json:"lastname"`
	PhoneNumbers []PhoneNumberRequest `json:"phoneNumbers"`
	Emails       []EmailRequest       `json:"emails"`
}

func (r NewContactRequest) Validate() *errs.AppError {

	if r.FirstName == utils.EmptyString &&
		r.LastName == utils.EmptyString &&
		len(r.PhoneNumbers) == 0 &&
		len(r.PhoneNumbers) == 0 {

		return errs.NewUnProcessableErr(errs.ErrUnprocessableRequest)

	}

	return nil
}
