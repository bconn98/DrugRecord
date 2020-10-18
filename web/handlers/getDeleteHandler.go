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

	"github.com/bconn98/DrugRecord/utils"
	"github.com/bconn98/DrugRecord/web/webUtils"
)

/**
Function: GetDeleteHandler
Description: Executes the delete template
*/
func GetDeleteHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR)
	}

	vars := mux.Vars(acRequest)

	lnId, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		utils.Log(err.Error(), utils.ERROR)
	}

	lasOrders := utils.GetOrder(lnId)

	webUtils.ExecuteTemplate(acWriter, "deleteSure.html", lasOrders[0])
}
