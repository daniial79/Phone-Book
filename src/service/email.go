package service

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
)

// EmailService Primary Port
type EmailService interface {
	AddNewEmails([]dto.AddEmailRequest, string) ([]dto.AddEmailResponse, *errs.AppError)
}
