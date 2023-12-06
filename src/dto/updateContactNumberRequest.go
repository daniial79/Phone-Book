package dto

type UpdateContactNumberRequest struct {
	NewPhoneNumber string `json:"newPhoneNumber,omitempty"`
	NewLabel       string `json:"newLabel,omitempty"`
}
