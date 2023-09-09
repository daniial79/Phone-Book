package core

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
)

// NumberRepository Number secondary port
type NumberRepository interface {
	AddNewNumber(Number) (*Number, *errs.AppError)
	CheckContactExistenceById(string) *errs.AppError
}

// Number core object definition
type Number struct {
	Id          string `db:"id"`
	ContactId   string `db:"contact_id"`
	PhoneNumber string `db:"number"`
	Label       string `db:"label"`
}

func (n Number) ToAddNumberResponseDto() *dto.AddNumberResponse {
	return &dto.AddNumberResponse{
		Id:        n.Id,
		ContactId: n.ContactId,
		Number:    n.PhoneNumber,
		Label:     n.Label,
	}
}
