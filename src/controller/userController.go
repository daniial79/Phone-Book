package controller

import (
	"errors"
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
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), utils.JsonIndentation)
	}

	response, appErr := c.service.SignUpUser(requestBody)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), utils.JsonIndentation)
	}

	accessToken, appErr := auth.NewAccessToken(requestBody.Username)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), utils.JsonIndentation)
	}
	auth.SetAuthCookie(&ctx, accessToken, utils.AccessTokenKey, utils.NewAccessTokenExpTime())

	refreshToken, appErr := auth.NewRefreshToken(requestBody.Username)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), utils.JsonIndentation)
	}
	auth.SetAuthCookie(&ctx, refreshToken, utils.RefreshTokenKey, utils.NewRefreshTokenExpTime())

	return ctx.JSONPretty(http.StatusCreated, response, utils.JsonIndentation)
}

func (c UserController) LogInController(ctx echo.Context) error {
	var requestBody dto.UserLoginRequest
	if err := ctx.Bind(&requestBody); err != nil {
		appErr := errs.NewUnProcessableErr(err.Error())
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), utils.JsonIndentation)
	}

	response, appErr := c.service.LogInUser(requestBody)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), utils.JsonIndentation)
	}

	accessToken, appErr := auth.NewAccessToken(requestBody.Username)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), utils.JsonIndentation)
	}
	auth.SetAuthCookie(&ctx, accessToken, utils.AccessTokenKey, utils.NewAccessTokenExpTime())

	refreshToken, appErr := auth.NewRefreshToken(requestBody.Username)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), utils.JsonIndentation)
	}
	auth.SetAuthCookie(&ctx, refreshToken, utils.RefreshTokenKey, utils.NewRefreshTokenExpTime())

	return ctx.JSONPretty(http.StatusOK, response, utils.JsonIndentation)
}

func (c UserController) RefreshTokenController(ctx echo.Context) error {
	cookie, err := ctx.Cookie(utils.RefreshTokenKey)
	if errors.Is(err, http.ErrNoCookie) {
		appErr := errs.NewUnAuthorizedErr(errs.CookieNotFoundErr)
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), utils.JsonIndentation)
	}

	refresherToken := cookie.Value
	username, appErr := auth.ParseJwtWithClaims(refresherToken)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), utils.JsonIndentation)
	}

	accessToken, appErr := auth.NewAccessToken(username)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), utils.JsonIndentation)
	}

	auth.SetAuthCookie(&ctx, accessToken, utils.AccessTokenKey, utils.NewAccessTokenExpTime())
	return ctx.JSONPretty(http.StatusOK, dto.NoContentResponse{}, utils.JsonIndentation)
}

func (c UserController) LogOutController(ctx echo.Context) error {
	appErr := auth.ExpireAuthCookies(&ctx)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), utils.JsonIndentation)
	}

	return ctx.JSON(http.StatusNoContent, dto.NoContentResponse{})
}
