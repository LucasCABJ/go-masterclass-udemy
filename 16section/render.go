package main

import (
	"net/http"
)

func (a *application) render(w http.ResponseWriter, filename string, data any) {
	if a.tp == nil {
		http.Error(w, "template renderer is nil", http.StatusInternalServerError)
	}
	a.tp.Render(w, filename, data)
}
