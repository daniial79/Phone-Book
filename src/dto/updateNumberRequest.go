package dto

type UpdateNumberRequest struct {
	NewPhoneNumber string `json:"newPhoneNumber,omitempty"`
	NewLabel       string `json:"newLabel,omitempty"`
}
