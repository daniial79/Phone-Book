package service

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
)

type UserService interface {
	SignUpUser(requestBody dto.CreateUserRequest) (*dto.CreateUserResponse, *errs.AppError)
	LogInUser(requestBody dto.UserLoginRequest) (*dto.UserLoginResponse, *errs.AppError)
}
