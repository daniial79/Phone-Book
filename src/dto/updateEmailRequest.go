package dto

import "github.com/daniial79/Phone-Book/src/errs"

type UpdateEmailRequest struct {
	NewAddress string `json:"newAddress"`
}

func (r *UpdateEmailRequest) Validate() *errs.AppError {
	emptyLiteralSample := UpdateEmailRequest{}
	requestLiteralValue := *r

	if requestLiteralValue == emptyLiteralSample {
		return errs.NewUnProcessableErr(errs.ErrUnprocessableRequest)
	}

	return nil
}
