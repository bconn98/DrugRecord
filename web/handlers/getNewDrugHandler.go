/**
File: getNewDrugHandler
Description: Gets new drug page
@author Bryan Conn
@date 1/4/2019
*/
package handlers

import (
	"net/http"

	"github.com/bconn98/DrugRecord/DrugRecord/mainUtils"
	"github.com/bconn98/DrugRecord/DrugRecord/web/utils"
)

/**
Function: GetNewDrugHandler
Description: Executes the new drug template
*/
func GetNewDrugHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.LogError(err.Error())
	}
	utils.ExecuteTemplate(acWriter, "newDrug.html", nil)
}
