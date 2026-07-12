package main

import (
	"net/http"
)

func (app *application) render(w http.ResponseWriter, filename string, data interface{}) {
	if app.tr == nil {
		http.Error(w, "template rendering engine not set", http.StatusInternalServerError)
		return
	}
	app.tr.render(w, filename, data)
}
