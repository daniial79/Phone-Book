package dto

// PhoneNumberRequest this phone number request is for initiating new contact
type PhoneNumberRequest struct {
	Number string `json:"number"`
	Label  string `json:"label"`
}

// EmailRequest this email request is for initiating new contact
type EmailRequest struct {
	Address string `json:"address"`
}

// NewContactRequest dto object definition
type NewContactRequest struct {
	FirstName    string               `json:"firstName"`
	LastName     string               `json:"lastname"`
	PhoneNumbers []PhoneNumberRequest `json:"phoneNumbers"`
	Emails       []EmailRequest       `json:"emails"`
}

func (r *NewContactRequest) IsValid() bool {
	requestLiteralValue := *r

	if requestLiteralValue.FirstName == "" &&
		requestLiteralValue.LastName == "" &&
		len(requestLiteralValue.PhoneNumbers) == 0 &&
		len(requestLiteralValue.PhoneNumbers) == 0 {

		return false

	}

	return true
}
