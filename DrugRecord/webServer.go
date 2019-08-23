/**
File: webServer
Description: Runs a database
@author Bryan Conn
@date 10/7/2018
*/
package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"./web/handlers"
	"./mainUtils"
)

/**
Function: main
Description: Creates a new web server at port 8080 and connects all of the handler functions
*/
func main() {
	var err error
	// If the file doesn't exist, create it, or append to the file
	mainUtils.GpcFile, err = os.OpenFile("c2.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	acRouter := mux.NewRouter()
	acRouter.HandleFunc("/", handlers.GetDatabaseHandler).Methods("GET")
	acRouter.HandleFunc("/home", handlers.GetHomeHandler).Methods("GET")
	acRouter.HandleFunc("/purchase", handlers.GetPurchaseHandler).Methods("GET")
	acRouter.HandleFunc("/prescription", handlers.GetPrescriptionHandler).Methods("GET")
	acRouter.HandleFunc("/audit", handlers.GetAuditHandler).Methods("GET")
	acRouter.HandleFunc("/login", handlers.GetLoginHandler).Methods("GET")
	acRouter.HandleFunc("/newDrug", handlers.GetNewDrugHandler).Methods("GET")
	acRouter.HandleFunc("/register", handlers.GetRegisterHandler).Methods("GET")
	acRouter.HandleFunc("/SignOut", handlers.GetSignOutHandler).Methods("GET")
	acRouter.HandleFunc("/closeWindow", handlers.GetCloseHandler).Methods("GET")
	acRouter.HandleFunc("/database", handlers.GetDatabaseHandler).Methods("GET")
	acRouter.HandleFunc("/edit", handlers.GetEditHandler).Methods("GET")
	acRouter.HandleFunc("/deleteSure", handlers.PostDeleteSureHandler).Methods("POST")
	acRouter.HandleFunc("/delete", handlers.GetDeleteHandler).Methods("GET")
	acRouter.HandleFunc("/delete", handlers.PostDeleteHandler).Methods("POST")
	acRouter.HandleFunc("/editQty", handlers.PostEditQtyHandler).Methods("POST")
	acRouter.HandleFunc("/edit", handlers.PostEditHandler).Methods("POST")
	acRouter.HandleFunc("/newDrug", handlers.PostNewDrugHandler).Methods("POST")
	acRouter.HandleFunc("/database", handlers.PostDatabaseHandler).Methods("POST")
	acRouter.HandleFunc("/login", handlers.PostLoginHandler).Methods("POST")
	acRouter.HandleFunc("/register", handlers.PostRegisterHandler).Methods("POST")
	acRouter.HandleFunc("/audit", handlers.PostAuditHandler).Methods("POST")
	acRouter.HandleFunc("/purchase", handlers.PostPurchaseHandler).Methods("POST")
	acRouter.HandleFunc("/prescription", handlers.PostPrescriptionHandler).Methods("POST")
	http.Handle("/web/assets/", http.StripPrefix("/web/assets", http.FileServer(http.Dir("./web/assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", acRouter)
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
