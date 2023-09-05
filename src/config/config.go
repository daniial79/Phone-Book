package config

import (
	"fmt"
	"github.com/daniial79/Phone-Book/src/logger"
	"log"
	"os"
)

var AppConf Config

func envVarSanityCheck(envVar string) string {
	if len(os.Getenv(envVar)) == 0 {
		log.Fatalln("Error you have to provide proper environment variable: " + envVar)
	}
	return os.Getenv(envVar)
}

func LoadConfig() {
	AppConf = Config{
		port:       envVarSanityCheck("PORT"),
		dbDriver:   envVarSanityCheck("DB_DRIVER"),
		dbHost:     envVarSanityCheck("DB_HOST"),
		dbPort:     envVarSanityCheck("DB_PORT"),
		dbUserName: envVarSanityCheck("DB_USERNAME"),
		dbPassword: envVarSanityCheck("DB_PASSWORD"),
		dbName:     envVarSanityCheck("DB_NAME"),
	}
	logger.Info("App config is successfully loaded")
}

type Config struct {
	port       string
	dbDriver   string
	dbHost     string
	dbPort     string
	dbUserName string
	dbPassword string
	dbName     string
}

func (c Config) GetPort() string {
	return ":" + c.port
}

func (c Config) GetDatabaseDriver() string {
	return c.dbDriver
}

func (c Config) GetDataSourceName() string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.dbHost,
		c.dbPort,
		c.dbUserName,
		c.dbPassword,
		c.dbName,
	)

	return dsn
}
