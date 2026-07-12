package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	userRepo    UserRepository
	templateDir string
	publicDir   string
	tr          *TemplateRenderer
}

func main() {
	db, err := connectDB("users_database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	templateDir := "./13web/templates"
	app := &application{
		errorLog:    log.New(os.Stderr, "ERROR\t", log.Ltime|log.LstdFlags|log.Lmicroseconds|log.Lshortfile),
		infoLog:     log.New(os.Stderr, "INFO\t", log.Ltime|log.LstdFlags),
		userRepo:    NewSqlUserRepository(db),
		templateDir: templateDir,
		publicDir:   filepath.Join(".", "13web", "public"),
		tr:          NewTemplateRenderer(templateDir, false),
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
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
