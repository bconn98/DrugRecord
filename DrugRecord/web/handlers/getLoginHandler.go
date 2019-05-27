/**
File: getLoginHandler
Description: Gets new login page
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
Function: GetLoginHandler
Description: Executes the login template
*/
func GetLoginHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	utils.ExecuteTemplate(acWriter, "login.html", nil)
}
