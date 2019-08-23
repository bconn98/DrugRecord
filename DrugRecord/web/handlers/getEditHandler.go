/**
File: getEditHandler
Description: Gets new audit page
@author Bryan Conn
@date 5/31/19
*/
package handlers

import (
	"net/http"

	"../../mainUtils"
	"../utils"
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
