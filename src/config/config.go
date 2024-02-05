package config

import (
	"fmt"
	"github.com/daniial79/Phone-Book/src/errs"
	"github.com/daniial79/Phone-Book/src/logger"
	"log"
	"os"
)

var appConf Config

func envVarSanityCheck(envVar string) string {
	if len(os.Getenv(envVar)) == 0 {
		log.Fatalln(errs.ErrMissedEnvVar + envVar)
	}
	return os.Getenv(envVar)
}

func LoadConfig() {
	appConf = Config{
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

func GetPort() string {
	return ":" + appConf.port
}

func GetDatabaseDriver() string {
	return appConf.dbDriver
}

func GetDataSourceName() string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		appConf.dbHost,
		appConf.dbPort,
		appConf.dbUserName,
		appConf.dbPassword,
		appConf.dbName,
		appConf.sslMode,
	)

	return dsn
}

func GetJwtKey() string {
	return appConf.jwtKey
}
