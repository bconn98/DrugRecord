/**
File: getDeleteHandler
Description: Gets new audit page
@author Bryan Conn
@date 6/2/19
*/
package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"../../mainUtils"
	"../utils"
)

/**
Function: GetDeleteHandler
Description: Executes the delete template
*/
func GetDeleteHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.LogError(err.Error())
	}

	vars := mux.Vars(acRequest)

	lnId, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		mainUtils.LogError(err.Error())
	}

	lasOrders := mainUtils.GetOrder(lnId)

	utils.ExecuteTemplate(acWriter, "deleteSure.html", lasOrders[0])
}
