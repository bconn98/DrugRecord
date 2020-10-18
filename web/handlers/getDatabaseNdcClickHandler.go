/**
File: getDatabaseHandler
Description: Gets a new database page
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/bconn98/DrugRecord/mainUtils"
	"github.com/bconn98/DrugRecord/web/utils"
)

/**
Function: GetDatabaseNdcHandler
Description: Executes the database template with the output data
*/
func GetDatabaseNdcClickHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	var lcErrorString string

	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.Log(err.Error(), mainUtils.ERROR)
	}
	vars := mux.Vars(acRequest)
	lcNdc := vars["ndc"]

	lcNdc, lcErrorString = utils.CheckNDC(lcNdc, lcErrorString)

	if lcErrorString != "" {
		utils.ExecuteTemplate(acWriter, "databaseDrug.html", nil)
		return
	}
	lcName, lcInput, lcForm, lcItemNum, lcSize, lcDate, lnQty, lasOrders := mainUtils.FindNDC(lcNdc)

	if lcName == "" && lcInput == "" && lcForm == "" && lcItemNum == "" {
		utils.ExecuteTemplate(acWriter, "databaseDrug.html", nil)
		return
	}

	lcDateString := lcDate.Month().String() + " " + strconv.Itoa(lcDate.Day()) + " " + strconv.Itoa(lcDate.Year())
	lsData := data{lcName, lcInput, lcForm, lcSize, lcDateString,
		lcItemNum, lnQty, lasOrders}
	utils.ExecuteTemplate(acWriter, "databaseDrug.html", lsData)
}
