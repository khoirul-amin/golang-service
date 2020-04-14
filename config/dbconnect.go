package config

import (
	"log"

	"database/sql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/simpleapp")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
