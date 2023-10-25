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
}
