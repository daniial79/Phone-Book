package dto

import "github.com/daniial79/Phone-Book/src/errs"

type UpdateNumberRequest struct {
	NewPhoneNumber string `json:"newPhoneNumber,omitempty"`
	NewLabel       string `json:"newLabel,omitempty"`
}

func (r *UpdateNumberRequest) Validate() *errs.AppError {
	zeroValuedSample := UpdateNumberRequest{}
	requestLiteralValue := *r

	if requestLiteralValue == zeroValuedSample {
		return errs.NewUnProcessableErr(errs.UnprocessableRequestErr)
	}

	return nil
}
