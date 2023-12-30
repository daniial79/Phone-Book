package dto

type CreateUserRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

func (r CreateUserRequest) IsValid() bool {
	if len(r.Password) < 8 {
		return false
	}

	return true
}
