/**
File: getEditHandler
Description: Gets new audit page
@author Bryan Conn
@date 5/31/19
*/
package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/bconn98/DrugRecord/DrugRecord/mainUtils"
	"github.com/bconn98/DrugRecord/DrugRecord/web/utils"
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

	vars := mux.Vars(acRequest)

	lnId, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		mainUtils.LogError(err.Error())
	}

	lasOrders := mainUtils.GetOrder(lnId)

	utils.ExecuteTemplate(acWriter, "editQty.html", lasOrders[0])
}
