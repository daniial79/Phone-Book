package controller

import (
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ContactController struct {
	cService service.ContactService
	nService service.NumberService
}

func NewContactController(
	cs service.ContactService,
	ns service.NumberService,
) ContactController {
	return ContactController{cService: cs, nService: ns}
}

func (c ContactController) NewContact(ctx echo.Context) error {
	var requestBody dto.NewContactRequest

	if err := ctx.Bind(&requestBody); err != nil {
		appErr := errs.NewUnProcessableErr(err.Error())
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	response, err := c.cService.NewContact(requestBody)
	if err != nil {
		return ctx.JSONPretty(err.StatusCode, err.AsMessage(), "  ")
	}

	return ctx.JSONPretty(http.StatusCreated, response, "  ")
}

func (c ContactController) AddNewNumber(ctx echo.Context) error {
	var requestBody dto.AddNumberRequest

	if err := ctx.Bind(&requestBody); err != nil {
		appErr := errs.NewUnProcessableErr(err.Error())
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	response, err := c.nService.AddNewNumbers(requestBody)
	if err != nil {
		return ctx.JSONPretty(err.StatusCode, err.AsMessage(), " ")
	}

	return ctx.JSONPretty(http.StatusCreated, response, " ")
}
