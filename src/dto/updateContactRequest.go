package dto

type UpdateContactRequest struct {
	NewFirstName string `json:"newFirstName,omitempty"`
	NewLastName  string `json:"newLastName,omitempty"`
}

func (r *UpdateContactRequest) IsValid() bool {
	zeroValuedSample := UpdateContactRequest{}
	requestLiteralValue := *r

	if requestLiteralValue == zeroValuedSample {
		return false
	}

	return true
}
