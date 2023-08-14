package app

import (
	"fmt"
	"github.com/daniial79/Phone-Book/src/config"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Start() {
	_ = config.NewConfig()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "hello from echo server",
		})
	})

	fmt.Println("server is up and running on port 8000...")
	if err := e.Start(":8000"); err != nil {
		log.Fatalln(err)
	}
}
