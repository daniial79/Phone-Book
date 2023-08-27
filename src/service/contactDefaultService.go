package service

import "github.com/daniial79/Phone-Book/src/core"

// ContactDefaultService primary actor
type ContactDefaultService struct {
	repo core.ContactRepository
}

func (s ContactDefaultService) NewContact() {

}
