package main
import (
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

func getHomeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "learning.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func postHomeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "learning.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getLoginHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getRegisterHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var templates = template.Must(template.ParseGlob("./templates/*.html"))

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getHomeHandler).Methods("GET")
	r.HandleFunc("/login", getLoginHandler).Methods("GET")
	r.HandleFunc("/register", getRegisterHandler).Methods("GET")
	r.HandleFunc("/", postHomeHandler).Methods("POST")
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}


