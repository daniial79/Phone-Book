package dto

type UpdateContactNumberResponse struct {
	UpdatedPhoneNumber string `json:"updatedPhoneNumber,omitempty"`
	UpdateLabel        string `json:"label,omitempty"`
}
