package core

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/google/uuid"
)

// Email core object definition
type Email struct {
	Id        uuid.UUID `db:"id"`
	ContactId uuid.UUID `db:"contact_id"`
	Address   string    `db:"address"`
}

func (e Email) ToAddEmailResponseDto() dto.AddEmailResponse {
	return dto.AddEmailResponse{
		Id:      e.Id,
		Address: e.Address,
	}
}

func (e Email) ToEmailResponseDto() dto.EmailResponse {
	return dto.EmailResponse(e)
}

func (e Email) ToUpdatedEmailResponseDto() dto.UpdateEmailResponse {
	return dto.UpdateEmailResponse{
		UpdatedAddress: e.Address,
	}
}
