package core

import (
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(*User) (*User, *errs.AppError)
}

type User struct {
	Id          uuid.UUID `db:"id"`
	Username    string
	Password    string
	PhoneNumber string `db:"phone_number"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}
