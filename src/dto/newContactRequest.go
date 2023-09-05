package dto

type PhoneNumberReq struct {
	Number string `json:"number"`
	Label  string `json:"label"`
}

type EmailReq struct {
	Address string `json:"address"`
}

// NewContactRequest dto object definition
type NewContactRequest struct {
	FirstName    string           `json:"firstName"`
	LastName     string           `json:"lastname"`
	PhoneNumbers []PhoneNumberReq `json:"phoneNumbers"`
	Emails       []EmailReq       `json:"emails"`
}
