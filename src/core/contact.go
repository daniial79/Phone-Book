package core

import "github.com/daniial79/Phone-Book/src/errs"

// Contact contact core object definition
type Contact struct {
	Id        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"LastName"`
}

// ContactRepository Contact secondary port
type ContactRepository interface {
	Create(Contact) (*Contact, *errs.AppError)
}
