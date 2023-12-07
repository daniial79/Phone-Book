package dto

type UpdateNumberResponse struct {
	UpdatedPhoneNumber string `json:"updatedPhoneNumber,omitempty"`
	UpdateLabel        string `json:"label,omitempty"`
}
