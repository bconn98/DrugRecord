/**
File: webServer
Description: Runs a database
@author Bryan Conn
@date 10/7/2018
*/
package main

import (
	"./web/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

/**
Function: main
Description: Creates a new web server at port 8080 and connects all of the handler functions
*/
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.GetDatabaseHandler).Methods("GET")
	r.HandleFunc("/home", handlers.GetHomeHandler).Methods("GET")
	r.HandleFunc("/purchase", handlers.GetPurchaseHandler).Methods("GET")
	r.HandleFunc("/prescription", handlers.GetPrescriptionHandler).Methods("GET")
	r.HandleFunc("/audit", handlers.GetAuditHandler).Methods("GET")
	r.HandleFunc("/login", handlers.GetLoginHandler).Methods("GET")
	r.HandleFunc("/newDrug", handlers.GetNewDrugHandler).Methods("GET")
	r.HandleFunc("/register", handlers.GetRegisterHandler).Methods("GET")
	r.HandleFunc("/signout", handlers.GetSignOutHandler).Methods("GET")
	r.HandleFunc("/closeWindow", handlers.GetCloseHandler).Methods("GET")
	r.HandleFunc("/database", handlers.GetDatabaseHandler).Methods("GET")
	r.HandleFunc("/newDrug", handlers.PostNewDrugHandler).Methods("POST")
	r.HandleFunc("/database", handlers.PostDatabaseHandler).Methods("POST")
	r.HandleFunc("/login", handlers.PostLoginHandler).Methods("POST")
	r.HandleFunc("/register", handlers.PostRegisterHandler).Methods("POST")
	r.HandleFunc("/audit", handlers.PostAuditHandler).Methods("POST")
	r.HandleFunc("/purchase", handlers.PostPurchaseHandler).Methods("POST")
	r.HandleFunc("/prescription", handlers.PostPrescriptionHandler).Methods("POST")
	http.Handle("/web/assets/", http.StripPrefix("/web/assets", http.FileServer(http.Dir("./web/assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", r)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
