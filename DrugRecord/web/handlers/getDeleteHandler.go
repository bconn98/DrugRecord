/**
File: getDeleteHandler
Description: Gets new audit page
@author Bryan Conn
@date 6/2/19
*/
package handlers

import (
	"log"
	"net/http"

	"../utils"
)

/**
Function: GetDeleteHandler
Description: Executes the delete template
*/
func GetDeleteHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	utils.ExecuteTemplate(acWriter, "delete.html", nil)
}
