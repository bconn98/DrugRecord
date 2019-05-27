/**
File: getSignOutHandler
Description: Gets new SignOut page
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"log"
	"net/http"

	"../utils"
)

/**
Function: GetSignOutHandler
Description: Executes the signOut template
*/
func GetSignOutHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	SetSignedOut()
	err := acRequest.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	utils.ExecuteTemplate(acWriter, "home.html", nil)
}
