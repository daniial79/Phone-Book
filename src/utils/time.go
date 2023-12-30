package utils

import (
	"fmt"
	"github.com/daniial79/Phone-Book/src/core"
	"time"
)

func GenCurrentDate(user *core.User) {
	year, monthInString, day := time.Now().Date()
	createdDate := fmt.Sprintf("%d-%d-%d", year, int(monthInString), day)
	user.CreatedAt, user.UpdatedAt = createdDate, createdDate
}
