package main
import (
	"net/http"
	"github.com/gorilla/mux"
	"./web/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.GetHomeHandler).Methods("GET")
	r.HandleFunc("/login", handlers.GetLoginHandler).Methods("GET")
	r.HandleFunc("/register", handlers.GetRegisterHandler).Methods("GET")
	r.HandleFunc("/signout", handlers.GetSignoutHandler).Methods("GET")
	r.HandleFunc("/database", handlers.GetDatabaseHandler).Methods("GET")
	r.HandleFunc("/database", handlers.PostDatabaseHandler).Methods("POST")
	r.HandleFunc("/login", handlers.PostLoginHandler).Methods("POST")
	r.HandleFunc("/register", handlers.PostRegisterHandler).Methods("POST")
	http.Handle("/web/assets/", http.StripPrefix("/web/assets", http.FileServer(http.Dir("./web/assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}