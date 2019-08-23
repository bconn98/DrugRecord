/**
File: getDatabaseHandler
Description: Gets a new database page
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"net/http"

	"../../mainUtils"
	"../utils"
)

var gbSignedIn = true // This should be false when there is a sign in feature

/**
Function: GetDatabaseHandler
Description: Executes the database template with the output data
*/
func GetDatabaseHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.LogError(err.Error())
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
