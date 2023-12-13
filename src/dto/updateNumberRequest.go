package dto

type UpdateNumberRequest struct {
	NewPhoneNumber string `json:"newPhoneNumber,omitempty"`
	NewLabel       string `json:"newLabel,omitempty"`
}

func (r *UpdateNumberRequest) IsValid() bool {
	zeroValuedSample := UpdateNumberRequest{}
	requestLiteralValue := *r

	if requestLiteralValue == zeroValuedSample {
		return false
	}

	return true
}
