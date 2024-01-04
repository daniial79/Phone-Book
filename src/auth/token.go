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

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, *errs.AppError) {
	claim := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(config.AppConf.GetJwtKey()))
	if err != nil {
		logger.Error("Error while generating auth-token for new user: " + err.Error())
		return "", errs.NewUnexpectedErr(errs.InternalErr)
	}

	return tokenString, nil
}

func SetAuthorizationCookie(ctx *echo.Context, authToken string) {
	cookie := new(http.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = authToken
	cookie.Expires = time.Now().Add(time.Minute * 30)
	cookie.Path = "/"
	cookie.HttpOnly = true

	(*ctx).SetCookie(cookie)
}
