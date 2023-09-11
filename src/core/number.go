package core

import (
	"github.com/daniial79/Phone-Book/src/dto"
)

// NumberRepository Number secondary port
type NumberRepository interface {
}

// Number core object definition
type Number struct {
	Id          string `db:"id"`
	ContactId   string `db:"contact_id"`
	PhoneNumber string `db:"number"`
	Label       string `db:"label"`
}

func (n Number) ToAddNumberResponseDto() dto.AddNumberResponse {
	return dto.AddNumberResponse{
		Id:     n.Id,
		Number: n.PhoneNumber,
		Label:  n.Label,
	}
}
