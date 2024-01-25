package mw

import (
	"errors"
	"github.com/daniial79/Phone-Book/src/auth"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/labstack/echo/v4"
	"net/http"
)

var CheckAccessToken echo.MiddlewareFunc = func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("Access-Token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				appErr := errs.NewBadRequestErr(errs.CookieNotFoundErr)
				return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
			}
		}

		username, appErr := auth.ParseJwtWithClaims(cookie.Value)
		if appErr != nil {
			return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
		}

		ctx.Set("username", username)
		return next(ctx)
	}
}
