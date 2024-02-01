package utils

import (
	"fmt"
	"time"
)

const (
	accessTokenExpMinute = 30
	refreshTokenExpHour  = 24
)

func NewCurrentDate() string {
	year, monthInString, day := time.Now().Date()
	return fmt.Sprintf("%d-%d-%d", year, int(monthInString), day)
}

func NewAccessTokenExpTime() time.Time {
	return time.Now().Add(time.Minute * accessTokenExpMinute)
}

func NewRefreshTokenExpTime() time.Time {
	return time.Now().Add(time.Hour * refreshTokenExpHour)
}
