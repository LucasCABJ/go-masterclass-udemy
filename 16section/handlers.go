package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, "index.html", nil)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")
		fmt.Println(email, password)
	}
	app.render(w, "login.html", nil)
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	app.render(w, "register.html", nil)
}

func (app *application) submit(w http.ResponseWriter, r *http.Request) {
	app.render(w, "submit.html", nil)
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, "contact.html", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, "about.html", nil)
}
