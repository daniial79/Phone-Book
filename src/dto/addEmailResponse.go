package dto

type AddEmailResponse struct {
	Id        string `json:"id"`
	ContactId string `json:"contact_id"`
	Address   string `json:"address"`
}
