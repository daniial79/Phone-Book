package controller

import (
	"fmt"
	"github.com/daniial79/Phone-Book/src/dto"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/logger"
	"github.com/daniial79/Phone-Book/src/service"
	"github.com/daniial79/Phone-Book/src/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ContactController struct {
	service service.ContactService
}

func NewContactController(
	s service.ContactService,
) ContactController {
	return ContactController{service: s}
}

func (c ContactController) NewContact(ctx echo.Context) error {
	var requestBody dto.NewContactRequest

	if err := ctx.Bind(&requestBody); err != nil {
		appErr := errs.NewUnProcessableErr(err.Error())
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	username, isThere := ctx.Get("username").(string)
	if !isThere {
		logger.Error("username is not embedded inside controller's context")
		appErr := errs.NewUnexpectedErr(errs.BadRequestErr)
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	response, err := c.service.NewContact(username, requestBody)
	if err != nil {
		return ctx.JSONPretty(err.StatusCode, err.AsMessage(), "  ")
	}

	return ctx.JSONPretty(http.StatusCreated, response, "  ")
}

func (c ContactController) AddNewNumbers(ctx echo.Context) error {
	contactIdAsString := ctx.Param("contactId")
	requestsBody := make([]dto.AddNumberRequest, 0)

	fmt.Printf("%s => %d\n", contactIdAsString, len(contactIdAsString))
	contactId, appErr := utils.ConvertStringToUUID([]byte(contactIdAsString))
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	if err := ctx.Bind(&requestsBody); err != nil {
		appErr = errs.NewUnProcessableErr(err.Error())
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	response, err := c.service.AddNewNumbers(requestsBody, contactId)
	if err != nil {
		return ctx.JSONPretty(err.StatusCode, err.AsMessage(), "  ")
	}

	return ctx.JSONPretty(http.StatusCreated, response, "  ")
}

func (c ContactController) AddNewEmails(ctx echo.Context) error {
	contactIdAsString := ctx.Param("contactId")
	request := make([]dto.AddEmailRequest, 0)

	contactId, appErr := utils.ConvertStringToUUID([]byte(contactIdAsString))
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	if err := ctx.Bind(&request); err != nil {
		appErr = errs.NewUnProcessableErr(err.Error())
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	response, err := c.service.AddNewEmails(request, contactId)
	if err != nil {
		return ctx.JSONPretty(err.StatusCode, err.AsMessage(), "  ")
	}

	return ctx.JSONPretty(http.StatusCreated, response, "  ")
}

func (c ContactController) GetContacts(ctx echo.Context) error {

	username, isThere := ctx.Get("username").(string)
	if !isThere {
		logger.Error("username is not embedded inside controller's context")
		appErr := errs.NewUnexpectedErr(errs.BadRequestErr)
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	response, appErr := c.service.GetContacts(username)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	return ctx.JSONPretty(http.StatusOK, response, "  ")
}

func (c ContactController) GetContactCredentials(ctx echo.Context) error {
	contactIdAsString := ctx.Param("contactId")

	contactId, appErr := utils.ConvertStringToUUID([]byte(contactIdAsString))
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	response, appErr := c.service.GetContactCredentials(contactId)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	return ctx.JSONPretty(http.StatusOK, response, "  ")
}

func (c ContactController) DeleteEmailFromContact(ctx echo.Context) error {
	contactIdAsString := ctx.Param("contactId")
	emailIdAsString := ctx.Param("emailId")

	contactId, appErr := utils.ConvertStringToUUID([]byte(contactIdAsString))
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	emailId, appErr := utils.ConvertStringToUUID([]byte(emailIdAsString))
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	response, appErr := c.service.DeleteEmailFromContact(contactId, emailId)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	return ctx.JSONPretty(http.StatusNoContent, response, "  ")
}

func (c ContactController) DeletePhoneNumberFromContact(ctx echo.Context) error {
	contactIdAsString := ctx.Param("contactId")
	phoneNumberIdAsString := ctx.Param("phoneNumberId")

	contactId, appErr := utils.ConvertStringToUUID([]byte(contactIdAsString))
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	phoneNumberId, appErr := utils.ConvertStringToUUID([]byte(phoneNumberIdAsString))
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	response, appErr := c.service.DeletePhoneNumberFromContact(contactId, phoneNumberId)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	return ctx.JSONPretty(http.StatusNoContent, response, "  ")
}

func (c ContactController) DeleteContact(ctx echo.Context) error {
	contactIdAsString := ctx.Param("contactId")

	contactId, appErr := utils.ConvertStringToUUID([]byte(contactIdAsString))
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	response, appErr := c.service.DeleteContact(contactId)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	return ctx.JSONPretty(http.StatusNoContent, response, "  ")
}

func (c ContactController) UpdatePhoneNumber(ctx echo.Context) error {
	contactIdAsString := ctx.Param("contactId")
	phoneNumberIdAsString := ctx.Param("phoneNumberId")

	contactId, appErr := utils.ConvertStringToUUID([]byte(contactIdAsString))
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	phoneNumberId, appErr := utils.ConvertStringToUUID([]byte(phoneNumberIdAsString))
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	var requestBody dto.UpdateNumberRequest
	if err := ctx.Bind(&requestBody); err != nil {
		appErr = errs.NewUnProcessableErr(err.Error())
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	response, appErr := c.service.UpdateContactNumber(contactId, phoneNumberId, requestBody)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), " ")
	}

	return ctx.JSONPretty(http.StatusOK, response, "  ")
}

func (c ContactController) UpdateEmail(ctx echo.Context) error {
	contactIdAsString := ctx.Param("contactId")
	emailIdAsString := ctx.Param("emailId")

	contactId, appErr := utils.ConvertStringToUUID([]byte(contactIdAsString))
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	emailId, appErr := utils.ConvertStringToUUID([]byte(emailIdAsString))
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	var requestBody dto.UpdateEmailRequest
	if err := ctx.Bind(&requestBody); err != nil {
		appErr = errs.NewUnProcessableErr(err.Error())
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	response, appErr := c.service.UpdateContactEmail(contactId, emailId, requestBody)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), " ")
	}

	return ctx.JSONPretty(http.StatusOK, response, "  ")
}

func (c ContactController) UpdateContact(ctx echo.Context) error {
	contactIdAsString := ctx.Param("contactId")

	contactId, appErr := utils.ConvertStringToUUID([]byte(contactIdAsString))
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	var requestBody dto.UpdateContactRequest
	if err := ctx.Bind(&requestBody); err != nil {
		appErr = errs.NewUnProcessableErr(err.Error())
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), "  ")
	}

	response, appErr := c.service.UpdateContact(contactId, requestBody)
	if appErr != nil {
		return ctx.JSONPretty(appErr.StatusCode, appErr.AsMessage(), " ")
	}

	return ctx.JSONPretty(http.StatusOK, response, "  ")
}
