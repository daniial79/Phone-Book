package service

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
)

// ContactService Primary port
type ContactService interface {
	NewContact(request dto.NewContactRequest) (*dto.ContactResponse, *errs.AppError)
}
