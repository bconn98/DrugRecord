/**
File: getRegisterHandler
Description: Gets new audit page
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"net/http"

	"../../mainUtils"
	"../utils"
)

/**
Function: GetRegisterHandler
Description: Execute the register template
*/
func GetRegisterHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.LogError(err.Error())
	}
	utils.ExecuteTemplate(acWriter, "register.html", nil)
}
