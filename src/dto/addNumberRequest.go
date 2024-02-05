package dto

import "github.com/daniial79/Phone-Book/src/errs"

type AddNumberRequest struct {
	Number string `json:"number"`
	Label  string `json:"label"`
}

func (r *AddNumberRequest) Validate() *errs.AppError {
	emptyLiteralSample := AddNumberRequest{}
	requestLiteralValue := *r

	if requestLiteralValue == emptyLiteralSample {
		return errs.NewUnexpectedErr(errs.ErrUnprocessableRequest)
	}

	return nil
}
