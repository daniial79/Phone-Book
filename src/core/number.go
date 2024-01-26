package core

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/google/uuid"
)

// Number core object definition
type Number struct {
	Id          uuid.UUID `db:"id"`
	ContactId   uuid.UUID `db:"contact_id"`
	PhoneNumber string    `db:"number"`
	Label       string    `db:"label"`
}

func (n Number) ToAddNumberResponseDto() dto.AddNumberResponse {
	return dto.AddNumberResponse{
		Id:     n.Id,
		Number: n.PhoneNumber,
		Label:  n.Label,
	}
}

func (n Number) ToPhoneNumberResponseDto() dto.PhoneNumberResponse {
	return dto.PhoneNumberResponse(n)
}

func (n Number) ToUpdatedNumberResponseDto() dto.UpdateNumberResponse {
	return dto.UpdateNumberResponse{
		UpdatedPhoneNumber: n.PhoneNumber,
		UpdateLabel:        n.Label,
	}
}
