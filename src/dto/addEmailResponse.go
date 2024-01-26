package dto

import "github.com/google/uuid"

type AddEmailResponse struct {
	Id      uuid.UUID `json:"id"`
	Address string    `json:"address"`
}
