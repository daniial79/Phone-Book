package dto

import "github.com/daniial79/Phone-Book/src/errs"

type UpdateContactRequest struct {
	NewFirstName string `json:"newFirstName,omitempty"`
	NewLastName  string `json:"newLastName,omitempty"`
}

func (r *UpdateContactRequest) Validate() *errs.AppError {
	zeroValuedSample := UpdateContactRequest{}
	requestLiteralValue := *r

	if requestLiteralValue == zeroValuedSample {
		return errs.NewUnProcessableErr(errs.ErrUnprocessableRequest)
	}

	return nil
}
