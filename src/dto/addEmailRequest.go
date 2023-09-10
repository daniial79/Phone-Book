package dto

type AddEmailRequest struct {
	ContactId string `json:"contact_id"`
	Address   string `json:"address"`
}
