package main

import (
	"net/http"
	"strconv"
)

const (
	loggedInUserKey = "logged_in_user_id"
)

func (app *application) readIntWithDefault(r *http.Request, key string, defaultValue int) int {
	v, err := strconv.Atoi(r.URL.Query().Get(key))
	if err != nil {
		return defaultValue
	}
	return v
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	filter := Filter{
		Query:    r.URL.Query().Get("q"),
		OrderBy:  r.URL.Query().Get("order_by"),
		Page:     app.readIntWithDefault(r, "page", 1),
		PageSize: app.readIntWithDefault(r, "page_size", 20),
	}

	posts, metadata, err := app.postRepository.GetAll(filter)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	app.render(w, r, "index.html", &templateData{
		Posts:    posts,
		Metadata: &metadata,
	})
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	if app.isAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

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
			IsEmail("email")

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
		app.session.Put(r, "flash", "You are logged in!")

		http.Redirect(w, r, "/submit", http.StatusSeeOther)
		return
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
			IsEmail("email")

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
		return
	}

	app.render(w, r, "register.html", &templateData{
		Form: NewForm(r.PostForm),
	})
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, loggedInUserKey)
	app.session.Put(r, "flash", "Logged out successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) vote(w http.ResponseWriter, r *http.Request) {
	postID := app.readIntWithDefault(r, "post_id", 0)
	if postID <= 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	userID := app.getUserFromContext(r.Context()).ID

	err := app.postRepository.AddVote(userID, postID)
	if err != nil {
		app.errorLog.Printf("Error adding vote: %v", err)
		app.session.Put(r, "flash", "voting failed")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	app.session.Put(r, "flash", "You voted successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) submit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		form := NewForm(r.PostForm)
		form.Required("title", "url").
			MaxLength("title", 255).
			MaxLength("url", 255).
			MinLength("title", 6).
			MinLength("url", 6)

		if !form.Valid() {
			form.Errors.Add("generic", "Invalid form data.")
			app.render(w, r, "submit.html", &templateData{
				Form: form,
			})
			return
		}

		title := r.FormValue("title")
		url := r.FormValue("url")
		userID := app.getUserFromContext(r.Context()).ID

		_, err := app.postRepository.CreatePost(title, url, userID)
		if err != nil {
			app.errorLog.Printf("Error creating post: %v", err)
			form.Errors.Add("generic", "An error occurred while creating the post. Please try again.")
			app.render(w, r, "submit.html", &templateData{Form: form})
			return
		}

		app.session.Put(r, "flash", "Post created successfully!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	app.render(w, r, "submit.html", &templateData{Form: NewForm(r.PostForm)})
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "contact.html", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "about.html", nil)
}
