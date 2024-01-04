package utils

import (
	"github.com/daniial79/Phone-Book/src/config"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Minute * 30).Unix(),
		},
	)

	tokenString, err := token.SignedString(config.AppConf.GetJwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
