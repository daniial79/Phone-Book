package auth

import (
	"errors"
	"github.com/daniial79/Phone-Book/src/config"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/logger"
	"github.com/daniial79/Phone-Book/src/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type UserClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func generateToken(
	username string,
	ExpiresAt time.Time,

) (string, *errs.AppError) {
	c := UserClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: ExpiresAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	stringToken, err := token.SignedString([]byte(config.GetJwtKey()))
	if err != nil {
		logger.Error("Error while generating string token")
		return "", errs.NewUnexpectedErr(errs.InternalErr)
	}

	return stringToken, nil
}

func NewAccessToken(username string) (string, *errs.AppError) {
	expirationTime := utils.NewAccessTokenExpTime()
	accessToken, err := generateToken(username, expirationTime)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func NewRefreshToken(username string) (string, *errs.AppError) {
	expirationTime := utils.NewRefreshTokenExpTime()
	refreshToken, err := generateToken(username, expirationTime)

	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func ParseJwtWithClaims(tokenString string) (string, *errs.AppError) {
	var uc UserClaim
	token, err := jwt.ParseWithClaims(tokenString, &uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJwtKey()), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return "", errs.NewUnAuthorizedErr(errs.UnauthorizedErr)
		}
		return "", errs.NewBadRequestErr(errs.BadRequestErr)
	}

	if !token.Valid {
		return "", errs.NewUnAuthorizedErr(errs.InvalidRefreshTokenErr)
	}

	return uc.Username, nil
}

func SetAuthCookie(
	ctx *echo.Context,
	stringToken string,
	name string,
	expAt time.Time,
) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = stringToken
	cookie.Expires = expAt
	cookie.Path = "/"
	cookie.HttpOnly = true

	(*ctx).SetCookie(cookie)
}
