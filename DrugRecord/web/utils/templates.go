package utils

import (
	"net/http"
	"html/template"
)

var templates *template.Template = template.Must(template.ParseGlob("./templates/*.html"))

func LoadTemplates(pattern string) {
	templates = template.Must(template.ParseGlob(pattern))	
}

func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}){
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}