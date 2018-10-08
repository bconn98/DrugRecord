/**
File: getAuditHandler
Description: Gets new audit page
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	"net/http"
	"../utils"
)

/**
Function: GetAuditHandler
Description: Executes the audit template
*/
func GetAuditHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "audit.html", nil)
}