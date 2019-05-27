/**
File: getRegisterHandler
Description: Gets new audit page
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
Function: GetRegisterHandler
Description: Execute the register template
*/
func GetRegisterHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	utils.ExecuteTemplate(acWriter, "register.html", nil)
}
