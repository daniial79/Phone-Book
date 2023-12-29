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

	g := r.Group("/api/v1/contacts")

	g.POST("/new", contactController.NewContact)
	g.POST("/:contactId/number", contactController.AddNewNumbers)
	g.POST("/:contactId/email", contactController.AddNewEmails)

	g.GET("/all", contactController.GetContacts)
	g.GET("/:contactId", contactController.GetContactCredentials)

	g.DELETE("/:contactId/emails/:emailId", contactController.DeleteEmailFromContact)
	g.DELETE("/:contactId/number/:phoneNumberId", contactController.DeletePhoneNumberFromContact)
	g.DELETE("/:contactId", contactController.DeleteContact)

	g.PATCH("/:contactId/number/:phoneNumberId", contactController.UpdatePhoneNumber)
	g.PATCH("/:contactId/emails/:emailId", contactController.UpdateEmail)
	g.PATCH("/:contactId", contactController.UpdateContact)
}
