package db

import (
	"database/sql"
	"github.com/daniial79/Phone-Book/src/config"
	"github.com/daniial79/Phone-Book/src/logger"
	_ "github.com/lib/pq"
)

func GetNewConnection() *sql.DB {
	db, err := sql.Open(
		config.GetDatabaseDriver(),
		config.GetDataSourceName(),
	)

	if err != nil {
		panic(err)
	}
	logger.Info("database DSN is valid")

	if err = db.Ping(); err != nil {
		panic(err)
	}
	logger.Info("Database connection is alive")

	return db
}
