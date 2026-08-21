package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func NewSQLiteConnection(filepath string) (*sql.DB, error) {
	dsn := filepath + "?_foreign_keys=on&_parseTime=true"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
