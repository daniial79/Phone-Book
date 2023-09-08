package dto

type PhoneNumberResponse struct {
	Id        string `json:"id"`
	ContactId string `json:"contactId"`
	Number    string `json:"number"`
	Label     string `json:"label"`
}
