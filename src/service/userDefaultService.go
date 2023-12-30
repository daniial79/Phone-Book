package service

import (
	"github.com/daniial79/Phone-Book/src/core"
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/logger"
	"github.com/daniial79/Phone-Book/src/utils"
)

type UserDefaultService struct {
	repo core.UserRepository
}

func NewUserDefaultService(repo core.UserRepository) UserDefaultService {
	return UserDefaultService{repo: repo}
}

func (s *UserDefaultService) CreateUser(requestBody dto.CreateUserRequest) (*dto.CreateUserResponse, *errs.AppError) {
	if !requestBody.IsValid() {
		return nil, errs.NewUnProcessableErr("Password should be at least 8 characters")
	}

	hashedPassword, err := utils.HashPassword(requestBody.Password)
	if err != nil {
		logger.Error("Error while hashing new user's password")
		return nil, errs.NewUnexpectedErr("Unexpected error happened")
	}

	coreTypedUser := core.User{
		Username:    requestBody.Username,
		Password:    hashedPassword,
		PhoneNumber: requestBody.PhoneNumber,
		CreatedAt:   utils.GenCurrentDate(),
		UpdatedAt:   utils.GenCurrentDate(),
	}

	createdUser, appErr := s.repo.CreateUser(coreTypedUser)
	if appErr != nil {
		return nil, appErr
	}

	response := createdUser.ToResponseDto()

	return response, nil
}
