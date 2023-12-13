package dto

type UpdateEmailRequest struct {
	NewAddress string `json:"newAddress"`
}

func (r *UpdateEmailRequest) IsValid() bool {
	emptyLiteralSample := UpdateEmailRequest{}
	requestLiteralValue := *r

	if requestLiteralValue == emptyLiteralSample {
		return false
	}

	return true
}
