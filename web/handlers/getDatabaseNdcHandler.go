/**
File: getDatabaseHandler
Description: Gets a new database page
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"net/http"

	"github.com/bconn98/DrugRecord/utils"
	"github.com/bconn98/DrugRecord/web/webUtils"
)

var gbSignedIn = true // This should be false when there is a sign in feature

/**
Function: GetDatabaseNdcHandler
Description: Executes the database template with the output data
*/
func GetDatabaseNdcHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR)
	}
	if !gbSignedIn {
		webUtils.ExecuteTemplate(acWriter, "home.html", "You are not signed in")
		return
	}
	webUtils.ExecuteTemplate(acWriter, "databaseDrug.html", nil)
}

func SetSignedIn() {
	gbSignedIn = true
}

func SetSignedOut() {
	gbSignedIn = false
}
