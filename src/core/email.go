package core

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
)

// EmailRepository secondary port
type EmailRepository interface {
	CheckContactExistenceById(string) *errs.AppError
	AddNewEmails([]Email) ([]Email, *errs.AppError)
}

// Email core object definition
type Email struct {
	Id        string `db:"id"`
	ContactId string `db:"contact_id"`
	Address   string `db:"address"`
}

func (e Email) ToAddEmailResponseDto() dto.AddEmailResponse {
	return dto.AddEmailResponse{
		Id:      e.Id,
		Address: e.Address,
	}
}
