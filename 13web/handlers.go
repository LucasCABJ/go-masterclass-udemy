package main

import (
	"fmt"
	"net/http"
)

var htmlContent = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>%s</title>
	</head>
	<body>
		%s
	<body>
	</html>
`

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid http method"))
		return
	}
	app.render(w, "index.html", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	content := fmt.Sprintf(htmlContent, "About", "<h1>About Page!</h1>")
	w.Write([]byte(content))
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	content := fmt.Sprintf(htmlContent, "Home", "<h1>Contact Page!</h1>")
	w.Write([]byte(content))
}
