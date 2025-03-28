package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error                           // Declare err separately to avoid shadowing DB
	DB, err = sql.Open("sqlite3", "api.db") // Assign to global DB

	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateTable()
}

func CreateTable() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	);`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		log.Fatal("Could not create events table:", err)
	}
}
