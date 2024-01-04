package dto

import "github.com/google/uuid"

type UserLoginResponse struct {
	Id uuid.UUID `json:"id"`
}
