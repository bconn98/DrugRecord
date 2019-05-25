/**
File: getSignOutHandler
Description: Gets new SignOut page
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"../utils"
	"log"
	"net/http"
)

/**
Function: GetSignOutHandler
Description: Executes the signOut template
*/
func GetSignOutHandler(w http.ResponseWriter, r *http.Request) {
	SetBad()
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	utils.ExecuteTemplate(w, "home.html", nil)
}
