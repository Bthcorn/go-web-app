package app

import (
	"html/template"
	"net/http"
)

type Templates struct {
	templates *template.Template
}

func NewTemplates() *Templates {
	// return a new instance of the Templates struct
	return &Templates{
		// reas all template from the templates folder
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
}

func (t *Templates) Render(w http.ResponseWriter, name string, data map[string]any) {
	t.templates.ExecuteTemplate(w, name, data)
}
