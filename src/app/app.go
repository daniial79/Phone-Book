package app

import (
	"github.com/daniial79/Phone-Book/src/config"
	"github.com/daniial79/Phone-Book/src/db"
	"github.com/daniial79/Phone-Book/src/logger"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Start() {
	config.LoadConfig()
	_ = db.GetNewConnection()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "hello from echo server",
		})
	})

	logger.Info("Server is up and running on port 8000...")
	if err := e.Start(config.AppConf.GetPort()); err != nil {
		log.Fatalln(err)
	}
}
