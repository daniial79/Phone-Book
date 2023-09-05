package errs

import "net/http"

type AppError struct {
	StatusCode int    `json:",omitempty"`
	Message    string `json:"message"`
}

func (a AppError) AsMessage() *AppError {
	return &AppError{
		Message: a.Message,
	}
}

func NewNotFoundErr(message string) *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

func NewUnexpectedErr(message string) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}

func NewValidationErr(message string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    message,
	}
}

func NewUnProcessableErr(message string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    message,
	}
}
