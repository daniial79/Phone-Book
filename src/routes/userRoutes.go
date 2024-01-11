package routes

import (
	"database/sql"
	"github.com/daniial79/Phone-Book/src/controller"
	"github.com/daniial79/Phone-Book/src/core"
	"github.com/daniial79/Phone-Book/src/service"
	"github.com/labstack/echo/v4"
)

func SetUserRoutes(e *echo.Echo, db *sql.DB) {
	userRepository := core.NewUserRepositoryDb(db)
	userService := service.NewUserDefaultService(userRepository)

	userController := controller.NewUserController(userService)

	r := e.Group("/api/v1")

	r.POST("/users/signup", userController.SignUpController)
	r.POST("/users/login", userController.LogInController)
	r.POST("/users/refresh", userController.RefreshTokenController)
}
