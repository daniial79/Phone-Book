package dto

type AddNumberRequest struct {
	Number string `json:"number"`
	Label  string `json:"label"`
}

func (r *AddNumberRequest) IsValid() bool {
	emptyLiteralSample := AddNumberRequest{}
	requestLiteralValue := *r

	if requestLiteralValue == emptyLiteralSample {
		return false
	}

	return true
}
