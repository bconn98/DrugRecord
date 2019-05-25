/**
File: getAuditHandler
Description: Gets new audit page
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
Function: GetAuditHandler
Description: Executes the audit template
*/
func GetAuditHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	utils.ExecuteTemplate(w, "audit.html", nil)
}
