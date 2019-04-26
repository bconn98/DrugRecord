/**
File: getAuditHandler
Description: Gets new audit page
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	. "../utils"
	"net/http"
)

/**
Function: GetCloseHandler
Description: Executes the close template
*/
func GetCloseHandler(w http.ResponseWriter, r *http.Request) {
	ExecuteTemplate(w, "closeWindow.html", nil)
	GetDatabaseHandler(w, r)
}