/**
File: getEditHandler
Description: Gets new audit page
@author Bryan Conn
@date 5/31/19
*/
package handlers

import (
	"log"
	"net/http"

	"../utils"
)

/**
Function: GetEditHandler
Description: Executes the edits template
*/
func GetEditHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	utils.ExecuteTemplate(acWriter, "edit.html", nil)
}
