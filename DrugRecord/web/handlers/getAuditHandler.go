/**
File: getAuditHandler
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
Function: GetAuditHandler
Description: Executes the audit template
*/
func GetAuditHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	utils.ExecuteTemplate(acWriter, "audit.html", nil)
}
