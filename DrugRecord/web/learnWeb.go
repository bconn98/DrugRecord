package main
import (
	"net/http"
	"html/template"
)

type Page struct {
	Title string
	Body  []byte
}


func viewHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "learning.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var templates = template.Must(template.ParseFiles("learning.html"))

func main() {
	http.HandleFunc("/", viewHandler)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}


