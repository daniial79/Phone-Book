package routes

import (
	"database/sql"
	"github.com/daniial79/Phone-Book/src/controller"
	"github.com/daniial79/Phone-Book/src/core"
	"github.com/daniial79/Phone-Book/src/service"
	"github.com/labstack/echo/v4"
)

func SetContactRoutes(r *echo.Echo, db *sql.DB) {
	contactRepository := core.NewContactRepositoryDb(db)
	contactService := service.NewContactDefaultService(contactRepository)

	contactController := controller.NewContactController(contactService)

	r.POST("/contacts", contactController.NewContact)
	r.POST("/contacts/:contactId/number", contactController.AddNewNumbers)
	r.POST("/contacts/:contactId/email", contactController.AddNewEmails)

	r.GET("/contacts", contactController.GetContacts)
	r.GET("/contacts/:contactId", contactController.GetContactCredentials)

	r.DELETE("/contacts/:contactId/emails/:emailId", contactController.DeleteEmailFromContact)
	r.DELETE("/contacts/:contactId/emails/:phoneNumberId", contactController.DeletePhoneNumberFromContact)
	r.DELETE("/contacts/:contactId", contactController.DeleteContact)

	r.PATCH("/contacts/:contactId/emails/:phoneNumberId", contactController.UpdatePhoneNumber)
	r.PATCH("/contacts/:contactId/emails/:emailId", contactController.UpdateEmail)
	r.PATCH("/contacts/:contactId", contactController.UpdateContact)
}
