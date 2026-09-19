package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	addr           string
	infoLog        *log.Logger
	errorLog       *log.Logger
	userRepository UserRepository
	templateDir    string
	tp             *TemplateRenderer
}

func main() {

	db, err := connectDB("users_database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &application{
		addr:           ":8080",
		errorLog:       log.New(os.Stderr, "ERROR\t", log.Ltime|log.LstdFlags|log.Lmicroseconds|log.Lshortfile),
		infoLog:        log.New(os.Stderr, "INFO\t", log.Ltime|log.LstdFlags),
		userRepository: NewSqlUserRepository(db),
		templateDir:    "./16section/templates",
	}
	app.tp = NewTemplateRenderer(app.templateDir, true)

	fmt.Println("Initializing server on port 8080")
	if err := app.serve(); err != nil {
		log.Fatal("Failed init server: %w", err)
	}
}

func connectDB(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to db: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database connection established")
	return db, nil
}
