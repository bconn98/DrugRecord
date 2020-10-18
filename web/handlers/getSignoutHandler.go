/**
File: getSignOutHandler
Description: Gets new SignOut page
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"net/http"

	"github.com/bconn98/DrugRecord/mainUtils"
	"github.com/bconn98/DrugRecord/web/utils"
)

/**
Function: GetSignOutHandler
Description: Executes the signOut template
*/
func GetSignOutHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	SetSignedOut()
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.Log(err.Error(), mainUtils.ERROR)
	}
	utils.ExecuteTemplate(acWriter, "home.html", nil)
}
