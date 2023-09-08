package routes

import (
	"database/sql"
	"github.com/daniial79/Phone-Book/src/controller"
	"github.com/daniial79/Phone-Book/src/core"
	"github.com/daniial79/Phone-Book/src/service"
	"github.com/labstack/echo/v4"
)

func SetContactRoutes(r *echo.Echo, db *sql.DB) {
	//wiring-up the contact section
	contactRepository := core.NewContactRepositoryDb(db)
	contactService := service.NewContactDefaultService(contactRepository)

	numberRepository := core.NewNumberRepositoryDb(db)
	numberService := service.NewNumberDefaultService(numberRepository)

	contactController := controller.NewContactController(contactService, numberService)

	r.POST("/contacts", contactController.NewContact)
	r.POST("/contacts/number", contactController.AddNewNumber)
}
