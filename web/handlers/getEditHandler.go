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

	utils "github.com/bconn98/DrugRecord/utils"
	webUtils "github.com/bconn98/DrugRecord/web/webUtils"
)

/**
Function: GetEditHandler
Description: Executes the edits template
*/
func GetEditHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
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

	webUtils.ExecuteTemplate(acWriter, "editQty.html", lasOrders[0])
}
