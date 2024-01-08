package auth

import (
	"github.com/daniial79/Phone-Book/src/config"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/logger"
	"github.com/daniial79/Phone-Book/src/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type userClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func generateToken(
	username string,
	ExpiresAt time.Time,

) (string, *errs.AppError) {
	c := userClaim{
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
