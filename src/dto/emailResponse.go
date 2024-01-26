package dto

import "github.com/google/uuid"

type EmailResponse struct {
	Id        uuid.UUID `json:"id"`
	ContactId uuid.UUID `json:"contactId"`
	Address   string    `json:"address"`
}
