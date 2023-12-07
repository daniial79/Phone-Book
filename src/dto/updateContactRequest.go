package dto

type UpdateContactRequest struct {
	NewFirstName string `json:"newFirstName,omitempty"`
	NewLastName  string `json:"newLastName,omitempty"`
}
