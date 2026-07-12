package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
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

	dbName := "users_database.db"

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

	createTable(db)

	lastId, err := createUser(db, "David", "david@david.com", "12345")
	if err != nil {
		log.Fatal(err)
	}

	lastId, err = createUser(db, "John", "john@john.com", "12345")
	if err != nil {
		log.Fatal(err)
	}

	lastId, err = createUser(db, "Kate", "kate@kate.com", "12345")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Last inserted: ", lastId)
}

func createTable(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("table created!")
}

func createUser(db *sql.DB, name, email, password string) (int64, error) {
	stmt := `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	result, err := db.Exec(stmt, name, email, hashPass)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}