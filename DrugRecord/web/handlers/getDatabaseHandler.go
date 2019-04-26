/**
File: getDatabaseHandler
Description: Gets a new database page
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	"net/http"
	"../utils"
)

var good = true	// This should be false when there is a sign in feature

/**
Function: GetDatabaseHandler
Description: Executes the database template with the output data
*/
func GetDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !good {
		utils.ExecuteTemplate(w, "home.html", "You are not signed in")
		return
	}
	utils.ExecuteTemplate(w,"database.html", nil)
}

func SetGood() {
	good = true
}

func SetBad() {
	good = false
}