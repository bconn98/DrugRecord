/**
File: getHomeHandler
Description: Gets new home page
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"net/http"

	"github.com/bconn98/DrugRecord/utils"
	"github.com/bconn98/DrugRecord/web/webUtils"
)

/**
Function: GetHomeHandler
Description: Executes the home template
*/
func GetHomeHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR)
	}
	webUtils.ExecuteTemplate(acWriter, "home.html", nil)
}
