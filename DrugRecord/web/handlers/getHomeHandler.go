/**
File: getHomeHandler
Description: Gets new home page
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"net/http"

	"github.com/bconn98/DrugRecord/DrugRecord/mainUtils"
	"github.com/bconn98/DrugRecord/DrugRecord/web/utils"
)

/**
Function: GetHomeHandler
Description: Executes the home template
*/
func GetHomeHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.LogError(err.Error())
	}
	utils.ExecuteTemplate(acWriter, "home.html", nil)
}
