package database

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)


func NewSQLiteConnection(filepath string) (*sql.DB, error){
	db, err := sql.Open("sqlite3",filepath);
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err := db.Ping();err != nil {
		return nil, err
	}
	return db, nil
}