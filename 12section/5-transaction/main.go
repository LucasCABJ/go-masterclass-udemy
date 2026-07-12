package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// database transaction
// ==========================

// 1. User creates an account
// 2. Create a wallet for the user
// 3. Want to top up the wallet for user
// 4. Write transaction log

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	hashed_password TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS profile (
	user_id INTEGER PRIMARY KEY REFERENCES users(user_id) ON DELETE CASCADE,
	avatar TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	Profile        Profile   `json:"profile"`
}

type Profile struct {
	UserID    int       `json:"user_id"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}

// prepared statement -> secure and optimized statement that can receive different parameters
func main() {

	dbName := "transaction_practice.db"

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatalf("Failed to connect to db: %s", err)
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

	initSchema(db)

	// userID, _ := createUserWithProfile(db, "Lucas Test", "tx@test.com", "12345", "http://testimage.com")
	// fmt.Println("Added user and profile for user_id:", userID)

	user, err := getUserByEmail(db, "tx@test.com")
	if err != nil {
		log.Fatal(err)
	}
	userMarshall, _ := json.MarshalIndent(user, "", " ")
	fmt.Printf("User found: %s\n", userMarshall)
}

func initSchema(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Schema initialized!")
}

// tsx need begin -> rollback or commit
func createUserWithProfile(db *sql.DB, name, email, password, avatar string) (int64, error) {
	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(name, email, hashPass)
	if err != nil {
		return 0, err
	}

	UserID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	profileStm, err := tx.PrepareContext(ctx, `INSERT INTO profile(user_id, avatar) VALUES (?, ?)`)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	defer profileStm.Close()

	_, err = profileStm.Exec(UserID, avatar)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, nil
	}

	return UserID, nil
}

func getUserByEmail(db *sql.DB, email string) (*User, error) {
	stmt := `SELECT id, name, email, hashed_password, u.created_at, p.user_id, p.avatar, p.created_at FROM users u INNER JOIN profile p ON u.id = p.user_id WHERE email = ?`
	row := db.QueryRow(stmt, email)
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt, &user.Profile.UserID, &user.Profile.Avatar, &user.Profile.CreatedAt)
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
