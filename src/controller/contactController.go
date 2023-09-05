package controller

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ContactController struct {
	service service.ContactService
}

func NewContactController(s service.ContactService) ContactController {
	return ContactController{service: s}
}

func (c ContactController) NewContact(ctx echo.Context) error {
	var requestBody dto.NewContactRequest

	if err := ctx.Bind(&requestBody); err != nil {
		appErr := errs.NewUnProcessableErr(err.Error())
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	response, err := c.service.NewContact(requestBody)
	if err != nil {
		return ctx.JSONPretty(err.StatusCode, err.AsMessage(), "  ")
	}

	return ctx.JSONPretty(http.StatusCreated, response, "  ")
}
