package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/golangcollege/sessions"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	addr           string
	infoLog        *log.Logger
	errorLog       *log.Logger
	userRepository UserRepository
	templateDir    string
	publicPath     string
	tp             *TemplateRenderer
	session        *sessions.Session
}

func main() {

	db, err := connectDB("users_database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	session := sessions.New([]byte("u46IpCV9y5VlurXXXODJEhgOY8m9JVE4"))
	session.Lifetime = 24 * time.Hour
	session.Secure = true
	session.SameSite = http.SameSiteLaxMode

	app := &application{
		addr:           ":8080",
		errorLog:       log.New(os.Stderr, "ERROR\t", log.Ltime|log.LstdFlags|log.Lmicroseconds|log.Lshortfile),
		infoLog:        log.New(os.Stderr, "INFO\t", log.Ltime|log.LstdFlags),
		userRepository: NewSqlUserRepository(db),
		templateDir:    "./templates",
		publicPath:     filepath.Join(".", "public"),
		session:        session,
	}
	app.tp = NewTemplateRenderer(app.templateDir, session, true)

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
