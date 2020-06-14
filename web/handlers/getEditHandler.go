/**
File: getEditHandler
Description: Gets new audit page
@author Bryan Conn
@date 5/31/19
*/
package handlers

import (
	"net/http"

	"github.com/bconn98/DrugRecord/mainUtils"
	"github.com/bconn98/DrugRecord/web/utils"
)

/**
Function: GetEditHandler
Description: Executes the edits template
*/
func GetEditHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.LogError(err.Error())
	}
	utils.ExecuteTemplate(acWriter, "edit.html", nil)
}
