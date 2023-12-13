package dto

type ContactCredentialsResponse struct {
	PhoneNumbers []PhoneNumberResponse `json:"phoneNumbers"`
	Emails       []EmailResponse       `json:"emails"`
}
