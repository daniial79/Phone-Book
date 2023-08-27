package core

import "github.com/daniial79/Phone-Book/src/errs"

// Email core object definition
type Email struct {
	ID        string `db:"id"`
	ContactId string `db:"contact_id"`
	Address   string `db:"address"`
}

// EmailRepository secondary port
type EmailRepository interface {
	Create(Email) (*Email, *errs.AppError)
}
