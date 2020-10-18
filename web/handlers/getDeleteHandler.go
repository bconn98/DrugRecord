package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/bconn98/DrugRecord/mainUtils"
	"github.com/bconn98/DrugRecord/web/utils"
)

/**
Function: GetDeleteHandler
Description: Executes the delete template
*/
func GetDeleteHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.Log(err.Error(), mainUtils.ERROR)
	}

	vars := mux.Vars(acRequest)

	lnId, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		mainUtils.Log(err.Error(), mainUtils.ERROR)
	}

	lasOrders := mainUtils.GetOrder(lnId)

	utils.ExecuteTemplate(acWriter, "delete.html", lasOrders[0])
}
