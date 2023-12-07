package dto

type UpdateContactResponse struct {
	UpdatedFirstName string `json:"updatedFirstName,omitempty"`
	UpdatedLastName  string `json:"updatedLastName,omitempty"`
}
