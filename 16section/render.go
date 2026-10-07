package main

import (
	"net/http"
)

func (a *application) render(w http.ResponseWriter, r *http.Request, filename string, data *templateData) {
	if a.tp == nil {
		http.Error(w, "template renderer is nil", http.StatusInternalServerError)
	}

	mergedData := a.buildDefaultTemplateData(data, r)
	a.tp.Render(w, filename, mergedData)
}

func (a *application) buildDefaultTemplateData(data *templateData, r *http.Request) *templateData {
	if data == nil {
		data = &templateData{}
	}
	data.Flash = a.session.PopString(r, "flash")
	return data
}
