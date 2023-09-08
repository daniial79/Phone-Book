package dto

type ContactResponse struct {
	Id           string                `json:"id"`
	FirstName    string                `json:"firstName"`
	LastName     string                `json:"lastName"`
	PhoneNumbers []PhoneNumberResponse `json:"phoneNumbers"`
	Emails       []EmailResponse       `json:"emails"`
}
