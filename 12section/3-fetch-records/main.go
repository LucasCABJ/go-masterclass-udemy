package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

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

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
}

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

	// lastId, err := createUser(db, "Jane", "jane@jane.com", "12345")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Last inserted: ", lastId)

	user, err := getUserByEmail(db, "jane@jane.com")
	if err != nil {
		log.Fatal(err)
	}

	bs, _ := json.MarshalIndent(user, "", " ")

	fmt.Printf("Found user: %v\n", string(bs))

	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	bs, _ = json.MarshalIndent(users, "", " ")

	fmt.Printf("Found users in DB: %v\n", string(bs))

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

func getUserByEmail(db *sql.DB, email string) (*User, error) {
	stmt := `SELECT id, name, email, hashed_password, created_at FROM users WHERE email = ?`
	row := db.QueryRow(stmt, email)
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func getUsers(db *sql.DB) ([]User, error) {
	var users []User
	stmt := `SELECT id, name, email, hashed_password, created_at FROM users`
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
