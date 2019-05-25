/**
File: getLoginHandler
Description: Gets new login page
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
Function: GetLoginHandler
Description: Executes the login template
*/
func GetLoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	utils.ExecuteTemplate(w, "login.html", nil)
}
