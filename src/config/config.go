package config

import (
	"fmt"
	"log"
	"os"
)

func envVarSanityCheck(envVar string) string {
	if len(os.Getenv(envVar)) == 0 {
		log.Fatalln("Error: you have to provide proper environment variable: " + envVar)
	}
	return os.Getenv(envVar)
}

func NewConfig() Config {
	return Config{
		port:       envVarSanityCheck("PORT"),
		dbDriver:   envVarSanityCheck("DB_DRIVER"),
		dbHost:     envVarSanityCheck("DB_HOST"),
		dbPort:     envVarSanityCheck("DB_PORT"),
		dbUserName: envVarSanityCheck("DB_USERNAME"),
		dbPassword: envVarSanityCheck("DB_PASSWORD"),
		dbName:     envVarSanityCheck("DB_NAME"),
	}
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

func (c Config) getPort() string {
	return ":" + c.port
}

func (c Config) getDatabaseDriver() string {
	return c.dbDriver
}

func (c Config) getDataSourceName() string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.dbHost,
		c.dbPort,
		c.dbUserName,
		c.dbPassword,
		c.dbName,
	)

	return dsn
}