package dto

import "github.com/google/uuid"

type PhoneNumberResponse struct {
	Id          uuid.UUID `json:"id"`
	ContactId   uuid.UUID `json:"contactId"`
	PhoneNumber string    `json:"number"`
	Label       string    `json:"label"`
}
