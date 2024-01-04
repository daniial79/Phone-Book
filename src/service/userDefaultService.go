package service

import (
	"errors"
	"github.com/daniial79/Phone-Book/src/auth"
	"github.com/daniial79/Phone-Book/src/core"
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/logger"
	"github.com/daniial79/Phone-Book/src/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserDefaultService struct {
	repo core.UserRepository
}

func NewUserDefaultService(repo core.UserRepository) UserDefaultService {
	return UserDefaultService{repo: repo}
}

func (s UserDefaultService) SignUpUser(requestBody dto.CreateUserRequest) (*dto.CreateUserResponse, *errs.AppError) {
	if appErr := requestBody.Validate(); appErr != nil {
		return nil, appErr
	}

	hashedPassword, err := auth.HashPassword(requestBody.Password)
	if err != nil {
		logger.Error("Error while hashing new user's password")
		return nil, errs.NewUnexpectedErr(errs.InternalErr)
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

	response := createdUser.ToCreatedResponseDto()

	return response, nil
}

func (s UserDefaultService) LogInUser(requestBody dto.UserLoginRequest) (*dto.UserLoginResponse, *errs.AppError) {
	if appErr := requestBody.Validate(); appErr != nil {
		return nil, appErr
	}

	username, password := requestBody.Username, requestBody.Password

	coreTypedUser, appErr := s.repo.GetUserByUsername(username)
	if appErr != nil {
		return nil, appErr
	}

	err := auth.ComparePasswords(coreTypedUser.Password, password)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, errs.NewUnAuthorizedErr(errs.MismatchedPasswords)
		}
		logger.Error("Error while comparing hash and plain text password: " + err.Error())
		return nil, errs.NewUnexpectedErr(errs.InternalErr)
	}

	response := coreTypedUser.ToLoggedInResponseDto()

	return response, nil

}
