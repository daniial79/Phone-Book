package dto

type AddEmailRequest struct {
	Address string `json:"address"`
}

func (r *AddEmailRequest) IsValid() bool {
	zeroValuedSample := AddEmailRequest{}
	requestLiteralValue := *r

	if requestLiteralValue == zeroValuedSample {
		return false
	}

	return true

}
