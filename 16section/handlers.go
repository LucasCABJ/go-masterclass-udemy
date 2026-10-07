package main

import (
	"net/http"
)

const (
	loggedInUserKey = "logged_in_user_id"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "index.html", nil)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Logged in: %s\n", app.session.GetString(r, loggedInUserKey))

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		form := NewForm(r.PostForm)
		form.Required("email", "password").
			MaxLength("email", 255).
			MaxLength("password", 255).
			MinLength("email", 6).
			MinLength("password", 3).
			Matches("email", EmailRX)

		if !form.Valid() {
			form.Errors.Add("generic", "Invalid form data.")
			app.render(w, r, "login.html", &templateData{
				Form: form,
			})
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		_, err := app.userRepository.Authenticate(email, password)
		if err != nil {
			form.Errors.Add("generic", err.Error())
			app.render(w, r, "login.html", &templateData{
				Form: form,
			})
			return
		}

		app.session.Put(r, loggedInUserKey, email)

		http.Redirect(w, r, "/submit", http.StatusSeeOther)
	}

	app.render(w, r, "login.html", &templateData{
		Form: NewForm(r.PostForm),
	})
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		form := NewForm(r.PostForm)
		form.Required("name", "email", "password", "avatar").
			MaxLength("name", 255).
			MaxLength("email", 255).
			MaxLength("password", 255).
			MaxLength("avatar", 255).
			MinLength("name", 6).
			MinLength("email", 6).
			MinLength("password", 5).
			MinLength("avatar", 10).
			Matches("email", EmailRX)

		if !form.Valid() {
			form.Errors.Add("generic", "Invalid form data.")
			app.render(w, r, "register.html", &templateData{
				Form: form,
			})
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		avatar := r.FormValue("avatar")

		_, err := app.userRepository.CreateUser(name, email, password, avatar)
		if err != nil {
			form.Errors.Add("generic", err.Error())
			app.render(w, r, "register.html", &templateData{
				Form: form,
			})
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	app.render(w, r, "register.html", &templateData{
		Form: NewForm(r.PostForm),
	})
}

func (app *application) submit(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "submit.html", &templateData{
		Form: NewForm(r.PostForm),
	})
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "contact.html", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "about.html", nil)
}
