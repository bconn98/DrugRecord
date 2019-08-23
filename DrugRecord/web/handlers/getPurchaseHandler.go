/**
File: getPurchaseHandler
Description: Gets new purchase page
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
Function: GetPurchaseHandler
Description: Executes the purchase template
*/
func GetPurchaseHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.LogError(err.Error())
	}
	utils.ExecuteTemplate(acWriter, "purchase.html", nil)
}
