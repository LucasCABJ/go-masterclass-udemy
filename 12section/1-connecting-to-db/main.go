package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	hashed_password BLOB NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

func main() {

	dbName := "data.db"
	// test db, in prod we shouldn't delete our db
	_ = os.Remove(dbName)

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatalf("Failed to connect to db")
	}
	defer func() {
		log.Printf("Closing DB")
		if err := db.Close(); err != nil {
			log.Printf("Failed to close DB")
		}
	}()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("database connection established")

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("table created!")

}
