package dto

import "github.com/daniial79/Phone-Book/src/errs"

type AddEmailRequest struct {
	Address string `json:"address"`
}

func (r *AddEmailRequest) Validate() *errs.AppError {
	zeroValuedSample := AddEmailRequest{}
	requestLiteralValue := *r

	if requestLiteralValue == zeroValuedSample {
		return errs.NewUnProcessableErr(errs.UnprocessableRequestErr)
	}

	return nil

}
