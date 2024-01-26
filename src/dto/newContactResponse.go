package dto

import "github.com/google/uuid"

type NewContactResponse struct {
	Id           uuid.UUID             `json:"id"`
	FirstName    string                `json:"firstName"`
	LastName     string                `json:"lastName"`
	PhoneNumbers []PhoneNumberResponse `json:"phoneNumbers,omitempty"`
	Emails       []EmailResponse       `json:"emails,omitempty"`
}
