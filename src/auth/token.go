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

func generateToken(username string, ExpiresAt time.Time) (string, *errs.AppError) {
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
		return utils.EmptyString, errs.NewUnexpectedErr(errs.ErrInternal)
	}

	return stringToken, nil
}

func NewAccessToken(username string) (string, *errs.AppError) {
	expirationTime := utils.NewAccessTokenExpTime()
	accessToken, err := generateToken(username, expirationTime)

	if err != nil {
		return utils.EmptyString, err
	}

	return accessToken, nil
}

func NewRefreshToken(username string) (string, *errs.AppError) {
	expirationTime := utils.NewRefreshTokenExpTime()
	refreshToken, err := generateToken(username, expirationTime)

	if err != nil {
		return utils.EmptyString, err
	}

	return refreshToken, nil
}

func ParseJwtWithClaims(tokenString string) (string, *errs.AppError) {
	var uc UserClaim
	token, err := jwt.ParseWithClaims(tokenString, &uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJwtKey()), nil
	})

	if err != nil {
		return utils.EmptyString, errs.NewUnAuthorizedErr(errs.ErrUnauthorized)
	}

	if !token.Valid {
		return utils.EmptyString, errs.NewUnAuthorizedErr(utils.InvalidToken)
	}

	return uc.Username, nil
}

func SetAuthCookie(ctx *echo.Context, stringToken string, name string, expAt time.Time) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = stringToken
	cookie.Expires = expAt
	cookie.Path = "/"
	cookie.HttpOnly = true

	(*ctx).SetCookie(cookie)
}

func ExpireAuthCookies(ctx *echo.Context) *errs.AppError {
	refreshCookie, err := (*ctx).Cookie(utils.RefreshTokenKey)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return errs.NewNotFoundErr(errs.ErrCookieNotFound)
		}
	}
	refreshCookie.Expires = time.Now()

	accessCookie, err := (*ctx).Cookie(utils.AccessTokenKey)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return errs.NewNotFoundErr(errs.ErrCookieNotFound)
		}
	}
	accessCookie.Expires = time.Now()

	return nil
}
