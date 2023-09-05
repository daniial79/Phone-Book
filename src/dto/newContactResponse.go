package dto

type PhoneNumberResp struct {
	Id        string `json:"id"`
	ContactId string `json:"contactId"`
	Number    string `json:"number"`
	Label     string `json:"label"`
}

type EmailResp struct {
	Id        string `json:"id"`
	ContactId string `json:"contactId"`
	Address   string `json:"address"`
}

type ContactResponse struct {
	Id           string            `json:"id"`
	FirstName    string            `json:"firstName"`
	LastName     string            `json:"lastName"`
	PhoneNumbers []PhoneNumberResp `json:"phoneNumbers"`
	Emails       []EmailResp       `json:"emails"`
}
