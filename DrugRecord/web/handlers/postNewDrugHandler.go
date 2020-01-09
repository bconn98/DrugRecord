/**
File: postNewDrugHandler
Description: Sends the new drug information
@author Bryan Conn
@date 1/4/2019
*/
package handlers

import (
	"net/http"

	"../../mainUtils"
	. "../utils"
)

/**
Function: PostNewDrugHelper
Description: Sends the purchase information to be added to the database
and executes the database template to refresh
*/
func PostNewDrugHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.LogError(err.Error())
	}

	var lcErrorString string
	lcId := acRequest.PostForm.Get("id")
	lcOldNdc := acRequest.PostForm.Get("oldNdc")
	lcNdc := acRequest.PostForm.Get("ndc")
	lcNdc, lcErrorString = CheckNDC(lcNdc, lcErrorString)
	lcName := acRequest.PostForm.Get("name")
	lcErrorString = CheckString(lcName, lcErrorString)
	lcForm := acRequest.PostForm.Get("form")
	lcItem := acRequest.PostForm.Get("item_num")
	lcPkgSize := acRequest.PostForm.Get("pkgSize")

	if lcErrorString != "" {

		if mainUtils.NewCheck(lcNdc) {
			mainUtils.UpdateOrderNdc(lcId, lcNdc)
			GetCloseHandler(acWriter, acRequest)
			return
		}

		ExecuteTemplate(acWriter, "newDrug.html", mainUtils.NewDrug{Error: lcErrorString, Ndc: lcNdc})
		return
	}
	mainUtils.UpdateDrug(lcPkgSize, lcForm, lcItem, lcName, lcNdc, 0, lcOldNdc)
	mainUtils.UpdateOrderNdc(lcId, lcNdc)
	GetCloseHandler(acWriter, acRequest)
	return
}
