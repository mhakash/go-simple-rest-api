package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var Database *sql.DB

func init() {
	var err error
	Database, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("could not connect to db %v", err)
	}

	statement, _ := Database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	_, err = statement.Exec()
	if err != nil {
		log.Fatalf("error occured %v", err)
	}
}
