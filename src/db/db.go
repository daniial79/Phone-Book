package db

import (
	"database/sql"
	"github.com/daniial79/Phone-Book/src/config"
	_ "github.com/lib/pq"
	"log"
)

func GetNewConnection() *sql.DB {
	db, err := sql.Open(
		config.AppConf.GetDatabaseDriver(),
		config.AppConf.GetDataSourceName(),
	)

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Fatalln("Error while closing connection: " + err.Error())
		}
	}(db)

	if err != nil {
		log.Fatalln("Error while trying to connect to database: ", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatalln("Pinging database failed: " + err.Error())
	}

	return db
}
