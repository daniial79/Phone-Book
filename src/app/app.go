package app

import (
	"fmt"
	"github.com/daniial79/Phone-Book/src/config"
	"github.com/daniial79/Phone-Book/src/db"
	"github.com/daniial79/Phone-Book/src/logger"
	"github.com/daniial79/Phone-Book/src/routes"
	"github.com/labstack/echo/v4"
	"log"
)

func Start() {
	config.LoadConfig()
	fmt.Println(config.AppConf.GetDataSourceName())
	dbClient := db.GetNewConnection()

	e := echo.New()

	routes.SetContactRoutes(e, dbClient)

	logger.Info("Server is up and running on port 8000...")
	if err := e.Start(config.AppConf.GetPort()); err != nil {
		log.Fatalln(err)
	}
}
