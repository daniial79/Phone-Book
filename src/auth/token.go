package auth

import (
	"github.com/daniial79/Phone-Book/src/config"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/logger"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type userClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewAccessToken(username string) (string, *errs.AppError) {
	c := userClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	accessToken, err := token.SignedString([]byte(config.GetJwtKey()))
	if err != nil {
		logger.Error("Error while generating access token for new user: " + err.Error())
		return "", errs.NewUnexpectedErr(errs.InternalErr)
	}

	return accessToken, nil
}

func NewRefreshToken(username string) (string, *errs.AppError) {
	c := userClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	refreshToken, err := token.SignedString([]byte(config.GetJwtKey()))
	if err != nil {
		logger.Error("Error while generating refresh token for new user: " + err.Error())
		return "", errs.NewUnexpectedErr(errs.InternalErr)
	}

	return refreshToken, nil
}

func SetAccessTokenCookie(ctx *echo.Context, accessToken string) {
	cookie := new(http.Cookie)
	cookie.Name = "Access-Token"
	cookie.Value = accessToken
	cookie.Expires = time.Now().Add(time.Minute * 30)
	cookie.Path = "/"
	cookie.HttpOnly = true

	(*ctx).SetCookie(cookie)
}

func SetRefreshTokenCookie(ctx *echo.Context, refreshToken string) {
	cookie := new(http.Cookie)
	cookie.Name = "Refresh-Token"
	cookie.Value = refreshToken
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.Path = "/"
	cookie.HttpOnly = true

	(*ctx).SetCookie(cookie)
}
