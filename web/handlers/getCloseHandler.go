/**
File: getAuditHandler
Description: Gets new audit page
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"net/http"

	. "github.com/bconn98/DrugRecord/DrugRecord/web/utils"
)

/**
Function: GetCloseHandler
Description: Executes the close template
*/
func GetCloseHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	ExecuteTemplate(acWriter, "closeWindow.html", nil)
	GetDatabaseNdcHandler(acWriter, acRequest)
}
