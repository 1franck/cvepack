package database

import (
	"cvepack/core/config"
	"cvepack/core/sqlite"
	"database/sql"
	"log"
)

func Connect() (db *sql.DB, closeDb func(db *sql.DB)) {
	db, err := sqlite.Connect(config.Default.DatabaseFilePath())
	if err != nil {
		log.Printf("error while connecting to database: %s", err)
		log.Fatal(err)
	}
	return db, func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("error while closing database: %s", err)
			log.Fatal(err)
		}
	}
}
