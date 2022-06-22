package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates inserts HTML templates into the variable templates
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// ExecuteTemplate renders a HTML page on the browser
func ExecuteTemplate(w http.ResponseWriter, templateFileName string, data interface{}) {
	templates.ExecuteTemplate(w, templateFileName, data)
}
