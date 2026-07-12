package main

import (
	"html/template"
	"net/http"
	"path"
	"path/filepath"
	"sync"
)

type TemplateRenderer struct {
	cache       map[string]*template.Template
	mutex       sync.RWMutex
	dev         bool
	templateDir string
}

func NewTemplateRenderer(templateDir string, devMode bool) *TemplateRenderer {
	return &TemplateRenderer{
		cache:       make(map[string]*template.Template),
		mutex:       sync.RWMutex{},
		dev:         devMode,
		templateDir: templateDir,
	}
}

func (t *TemplateRenderer) render(w http.ResponseWriter, templateName string, data interface{}) {
	tmpl, err := t.getTemplate(templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TemplateRenderer) getTemplate(templateName string) (*template.Template, error) {
	if !t.dev {
		t.mutex.RLock()
		if tmpl, ok := t.cache[templateName]; ok {
			t.mutex.RUnlock()
			return tmpl, nil
		}
		t.mutex.RUnlock()
	}

	tmpl, err := t.parseTemplate(templateName)
	if err != nil {
		return nil, err
	}

	if !t.dev {
		t.mutex.Lock()
		t.cache[templateName] = tmpl
		t.mutex.Unlock()
	}

	return tmpl, nil
}

func (t *TemplateRenderer) parseTemplate(templateName string) (*template.Template, error) {
	templatePath := path.Join(t.templateDir, templateName)
	files := []string{templatePath}
	layoutPath := path.Join(t.templateDir, "layouts/*.html")
	layouts, err := filepath.Glob(layoutPath)
	if err == nil {
		files = append(files, layouts...)
	}
	partialsPath := path.Join(t.templateDir, "partials/*.html")
	partials, err := filepath.Glob(partialsPath)
	if err == nil {
		files = append(files, partials...)
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
