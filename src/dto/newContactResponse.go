package dto

type NewContactResponse struct {
	Id           string                `json:"id"`
	FirstName    string                `json:"firstName"`
	LastName     string                `json:"lastName"`
	PhoneNumbers []PhoneNumberResponse `json:"phoneNumbers,omitempty"`
	Emails       []EmailResponse       `json:"emails,omitempty"`
}
