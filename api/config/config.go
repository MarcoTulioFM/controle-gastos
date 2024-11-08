package config

import (
	"database/sql"
	"log"
)

func Init() (*sql.DB, error) {
	var err error

	db, err := initializePostgres()
	if err != nil{
		log.Fatal(err)
	}
	return db, nil

}