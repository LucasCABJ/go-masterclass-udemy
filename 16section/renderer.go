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
	devMode     bool
	templateDir string
}

func NewTemplateRenderer(templateDir string, isDev bool) *TemplateRenderer {
	return &TemplateRenderer{
		cache:       make(map[string]*template.Template),
		devMode:     isDev,
		templateDir: templateDir,
	}
}

func (t *TemplateRenderer) Render(w http.ResponseWriter, templateName string, data any) {
	tmpl, err := t.getTemplate(templateName)
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

func (t *TemplateRenderer) getTemplate(name string) (*template.Template, error) {
	if !t.devMode {
		t.mutex.RLock()
		if tmpl, ok := t.cache[name]; ok {
			t.mutex.RUnlock()
			return tmpl, nil
		}
		t.mutex.RUnlock()
	}

	tmpl, err := t.parseTemplate(name)
	if err != nil {
		return nil, err
	}

	if !t.devMode {
		t.mutex.Lock()
		t.cache[name] = tmpl
		t.mutex.Unlock()
	}

	return tmpl, nil
}

func (t *TemplateRenderer) parseTemplate(name string) (*template.Template, error) {
	templatePath := path.Join(t.templateDir, name)

	files := []string{templatePath}

	layoutPath := path.Join(t.templateDir, "layout/*.html")
	layouts, err := filepath.Glob(layoutPath)
	if err == nil {
		files = append(files, layouts...)
	}

	partialPath := path.Join(t.templateDir, "partials/*.html")
	partials, err := filepath.Glob(partialPath)
	if err == nil {
		files = append(files, partials...)
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}

	return tmpl, err
}
