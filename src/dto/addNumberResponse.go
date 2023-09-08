package dto

type AddNumberResponse struct {
	Id        string `json:"id"`
	ContactId string `json:"contact_id"`
	Number    string `json:"number"`
	Label     string `json:"label"`
}
