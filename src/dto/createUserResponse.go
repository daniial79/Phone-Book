package dto

import "github.com/google/uuid"

type CreateUserResponse struct {
	Id uuid.UUID `json:"id"`
}
