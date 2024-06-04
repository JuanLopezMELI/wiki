package utils

import (
	"html/template"
	"net/http"
	"web-app/entities"
)

var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html"))

func RenderTemplate(templateName string, w http.ResponseWriter, p *entities.Page) {
	err := templates.ExecuteTemplate(w, templateName+".html", p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
