/**
File: getDatabaseHandler
Description: Gets a new database page
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"../utils"
	"log"
	"net/http"
)

var gbSignedIn = true // This should be false when there is a sign in feature

/**
Function: GetDatabaseHandler
Description: Executes the database template with the output data
*/
func GetDatabaseHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	if !gbSignedIn {
		utils.ExecuteTemplate(acWriter, "home.html", "You are not signed in")
		return
	}
	utils.ExecuteTemplate(acWriter, "database.html", nil)
}

func SetSignedIn() {
	gbSignedIn = true
}

func SetSignedOut() {
	gbSignedIn = false
}
