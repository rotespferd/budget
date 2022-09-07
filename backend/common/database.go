package common

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./db/budget.sqlite")

	if err != nil {
		log.Panic(err)
	}

	return db
}
