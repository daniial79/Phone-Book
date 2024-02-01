package dto

import (
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/utils"
)

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r UserLoginRequest) Validate() *errs.AppError {
	if r.Username == utils.EmptyString && r.Password == utils.EmptyString {
		return errs.NewUnProcessableErr(errs.UnprocessableRequestErr)
	}

	if r.Username == utils.EmptyString || r.Password == utils.EmptyString {
		return errs.NewCredentialsErr(errs.InsufficientCredentialsErr)
	}

	if len(r.Password) < 8 {
		return errs.NewUnProcessableErr(errs.ShortPasswordErr)
	}

	return nil
}
