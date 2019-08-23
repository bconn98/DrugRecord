/**
File: getLoginHandler
Description: Gets new login page
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
Function: GetLoginHandler
Description: Executes the login template
*/
func GetLoginHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.LogError(err.Error())
	}
	utils.ExecuteTemplate(acWriter, "login.html", nil)
}
