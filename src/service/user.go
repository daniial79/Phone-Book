package service

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
)

type UserService interface {
	CreateUser(requestBody dto.CreateUserRequest) (*dto.CreateUserResponse, *errs.AppError)
}
