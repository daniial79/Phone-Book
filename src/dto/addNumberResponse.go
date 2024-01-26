package dto

import "github.com/google/uuid"

type AddNumberResponse struct {
	Id     uuid.UUID `json:"id"`
	Number string    `json:"number"`
	Label  string    `json:"label"`
}
