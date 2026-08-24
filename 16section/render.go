package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func (a *application) render(w http.ResponseWriter, filename string, data any) {
	fullPath := path.Join(a.templateDir, filename)
	fmt.Println(fullPath)
	tmpl, err := template.ParseFiles(fullPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
