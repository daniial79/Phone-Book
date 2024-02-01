package core

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
)

type UserRepository interface {
	CreateUser(User) (*User, *errs.AppError)
	GetUserByUsername(username string) (*User, *errs.AppError)
}

type User struct {
	Id          string `db:"id"`
	Username    string
	Password    string
	PhoneNumber string `db:"phone_number"`
	Role        string `db:"user_role"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

func (u User) ToCreatedResponseDto() *dto.CreateUserResponse {
	return &dto.CreateUserResponse{
		Id: u.Id,
	}
}

func (u User) ToLoggedInResponseDto() *dto.UserLoginResponse {
	return &dto.UserLoginResponse{
		Id: u.Id,
	}
}
