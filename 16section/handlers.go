package main

import (
	"fmt"
	"net/http"
)

var htmlContentTemplate = `
	<!DOCTYPE html>
	<html>
		<head>
			<title>%s</title>
		</head>
		<body>
			%s
		</body>
	</html>
`

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// homeContent := fmt.Sprintf(htmlContentTemplate, "Home", "<h1>This is home page!</h1>")
	app.render(w, "index.html", nil)
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	contactContent := fmt.Sprintf(htmlContentTemplate, "Contact", "<h1>This is contact page!</h1>")
	writeHtmlResponse(w, contactContent)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	aboutContent := fmt.Sprintf(htmlContentTemplate, "About", "<h1>This is about page!</h1>")
	writeHtmlResponse(w, aboutContent)
}

func writeHtmlResponse(w http.ResponseWriter, html string) {
	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte(html))
}
