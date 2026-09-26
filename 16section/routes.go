package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	defaultMiddlewares := alice.New(app.logger, app.panicRecover)
	secureMiddleware := alice.New(app.session.Enable)

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(app.publicPath))))

	mux.Handle("/", secureMiddleware.ThenFunc(app.home))
	mux.Handle("/login", secureMiddleware.ThenFunc(app.login))
	mux.Handle("/submit", secureMiddleware.Append(app.requireAuth).ThenFunc(app.login))
	mux.Handle("/register", secureMiddleware.ThenFunc(app.register))
	mux.HandleFunc("/about", app.about)
	mux.HandleFunc("/contact", app.contact)

	return defaultMiddlewares.Then(mux)
}
