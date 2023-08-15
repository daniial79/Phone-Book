package db

import (
	"database/sql"
	"github.com/daniial79/Phone-Book/src/config"
	_ "github.com/lib/pq"
)

func GetNewConnection() *sql.DB {
	db, err := sql.Open(
		config.AppConf.GetDatabaseDriver(),
		config.AppConf.GetDataSourceName(),
	)

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}
