package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"learning_go_udemy/12section/6-repository/repository"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// prepared statement -> secure and optimized statement that can receive different parameters
func main() {

	dbName := "transaction_practice.db"
	db, err := connectDB(dbName)
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewSqlUserRepository(db)

	// userID, _ := createUserWithProfile(db, "Lucas Test", "tx@test.com", "12345", "http://testimage.com")
	// fmt.Println("Added user and profile for user_id:", userID)

	user, err := repository.GetUserByEmail("tx@test.com")
	if err != nil {
		log.Fatal(err)
	}
	userMarshall, _ := json.MarshalIndent(user, "", " ")
	fmt.Printf("User found: %s\n", userMarshall)
}

func connectDB(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to db: %s", err)
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
	return db, nil
}