/**
File: getDatabaseHandler
Description: Gets a new database page
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jimlawless/whereami"

	"github.com/bconn98/DrugRecord/utils"
	"github.com/bconn98/DrugRecord/web/webUtils"
)

/**
Function: GetDatabaseNdcHandler
Description: Executes the database template with the output data
*/
func GetDatabaseNdcClickHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	var lcErrorString string

	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}
	vars := mux.Vars(acRequest)
	lcNdc := vars["ndc"]

	lcNdc, lcErrorString = webUtils.CheckNDC(lcNdc, lcErrorString)

	if lcErrorString != "" {
		webUtils.ExecuteTemplate(acWriter, "databaseDrug.html", nil)
		return
	}
	lcName, lcInput, lcForm, lcItemNum, lcSize, lcDate, lnQty, lasOrders := utils.FindNDC(lcNdc)

	if lcName == "" && lcInput == "" && lcForm == "" && lcItemNum == "" {
		webUtils.ExecuteTemplate(acWriter, "databaseDrug.html", nil)
		return
	}

	lcDateString := lcDate.Month().String() + " " + strconv.Itoa(lcDate.Day()) + " " + strconv.Itoa(lcDate.Year())
	lsData := data{lcName, lcInput, lcForm, lcSize, lcDateString,
		lcItemNum, lnQty, lasOrders}
	webUtils.ExecuteTemplate(acWriter, "databaseDrug.html", lsData)
}
