package config

import (
	"fmt"
	"github.com/daniial79/Phone-Book/src/logger"
	"log"
	"os"
)

const (
	errMissedEnvVar = "Error you have to provide proper environment variable: "
)

var AppConf Config

func envVarSanityCheck(envVar string) string {
	if len(os.Getenv(envVar)) == 0 {
		log.Fatalln(errMissedEnvVar + envVar)
	}
	return os.Getenv(envVar)
}

func LoadConfig() {
	AppConf = Config{
		port:       envVarSanityCheck("PORT"),
		jwtKey:     envVarSanityCheck("JWT_KEY"),
		dbDriver:   envVarSanityCheck("DB_DRIVER"),
		dbHost:     envVarSanityCheck("DB_HOST"),
		dbPort:     envVarSanityCheck("DB_PORT"),
		dbUserName: envVarSanityCheck("DB_USERNAME"),
		dbPassword: envVarSanityCheck("DB_PASSWORD"),
		dbName:     envVarSanityCheck("DB_NAME"),
		sslMode:    envVarSanityCheck("SSL_MODE"),
	}
	logger.Info("App config is successfully loaded")
}

type Config struct {
	port       string
	jwtKey     string
	dbDriver   string
	dbHost     string
	dbPort     string
	dbUserName string
	dbPassword string
	dbName     string
	sslMode    string
}

func (c Config) GetPort() string {
	return ":" + c.port
}

func (c Config) GetDatabaseDriver() string {
	return c.dbDriver
}

func (c Config) GetDataSourceName() string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.dbHost,
		c.dbPort,
		c.dbUserName,
		c.dbPassword,
		c.dbName,
		c.sslMode,
	)

	return dsn
}

func (c Config) GetJwtKey() string {
	return c.jwtKey
}
