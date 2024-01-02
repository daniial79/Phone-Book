package dto

import "github.com/daniial79/Phone-Book/src/errs"

type CreateUserRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

func (r CreateUserRequest) Validate() *errs.AppError {
	if r.Username == "" &&
		r.Password == "" &&
		r.PhoneNumber == "" {
		return errs.NewUnProcessableErr(errs.UnprocessableRequestErr)
	}

	if len(r.Password) < 8 {
		return errs.NewUnProcessableErr(errs.ShortPasswordErr)
	}

	return nil
}
