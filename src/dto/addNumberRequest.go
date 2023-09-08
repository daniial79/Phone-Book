package dto

type AddNumberRequest struct {
	ContactId string `json:"contact_id"`
	Number    string `json:"number"`
	Label     string `json:"label"`
}
