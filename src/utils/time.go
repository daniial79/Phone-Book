package utils

import (
	"fmt"
	"time"
)

func GenCurrentDate() string {
	year, monthInString, day := time.Now().Date()
	return fmt.Sprintf("%d-%d-%d", year, int(monthInString), day)
}
