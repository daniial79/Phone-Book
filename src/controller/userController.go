package controller

import (
	"github.com/daniial79/Phone-Book/src/auth"
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/service"
	"github.com/daniial79/Phone-Book/src/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	service service.UserService
}

func NewUserController(s service.UserService) UserController {
	return UserController{
		service: s,
	}
}

func (c UserController) SignUpController(ctx echo.Context) error {
	var requestBody dto.CreateUserRequest
	if err := ctx.Bind(&requestBody); err != nil {
		appErr := errs.NewUnProcessableErr(err.Error())
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	response, appErr := c.service.SignUpUser(requestBody)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	accessToken, appErr := auth.NewAccessToken(requestBody.Username)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}
	auth.SetAuthCookie(&ctx, accessToken, "Access-Token", utils.NewAccessTokenExpTime())

	refreshToken, appErr := auth.NewRefreshToken(requestBody.Username)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}
	auth.SetAuthCookie(&ctx, refreshToken, "Refresh-Token", utils.NewRefreshTokenExpTime())

	return ctx.JSONPretty(http.StatusCreated, response, "  ")
}

func (c UserController) LogInController(ctx echo.Context) error {
	var requestBody dto.UserLoginRequest
	if err := ctx.Bind(&requestBody); err != nil {
		appErr := errs.NewUnProcessableErr(err.Error())
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	response, appErr := c.service.LogInUser(requestBody)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	accessToken, appErr := auth.NewAccessToken(requestBody.Username)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}
	auth.SetAuthCookie(&ctx, accessToken, "Access-Token", utils.NewAccessTokenExpTime())

	refreshToken, appErr := auth.NewRefreshToken(requestBody.Username)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}
	auth.SetAuthCookie(&ctx, refreshToken, "Refresh-Token", utils.NewRefreshTokenExpTime())

	return ctx.JSONPretty(http.StatusOK, response, "  ")
}
