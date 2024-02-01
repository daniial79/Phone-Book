package dto

import (
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/utils"
)

type CreateUserRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role,omitempty"`
}

func (r CreateUserRequest) Validate() *errs.AppError {
	if r.Username == utils.EmptyString &&
		r.Password == utils.EmptyString &&
		r.PhoneNumber == utils.EmptyString {
		return errs.NewUnProcessableErr(errs.UnprocessableRequestErr)
	}

	if r.Username == utils.EmptyString ||
		r.Password == utils.EmptyString ||
		r.PhoneNumber == utils.EmptyString {
		return errs.NewCredentialsErr(errs.InsufficientCredentialsErr)
	}

	if len(r.Password) < utils.MinimumPasswordLength {
		return errs.NewUnProcessableErr(errs.ShortPasswordErr)
	}

	if r.Role == utils.EmptyString || r.Role == utils.UserRoleString {
		return nil
	}

	if r.Role != utils.AdminRoleString {
		return errs.NewUnProcessableErr(errs.UserRoleErr)
	}

	return nil
}

func (r CreateUserRequest) SetUserRole() string {
	if r.Role != utils.AdminRoleString {
		return utils.UserRoleString
	}

	return utils.AdminRoleString
}
