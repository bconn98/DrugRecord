/**
File: postNewDrugHandler
Description: Sends the new drug information
@author Bryan Conn
@date 1/4/2019
*/
package handlers

import (
	"log"
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
		log.Fatal(err)
	}
	var lcErrorString string
	lcNdc := acRequest.PostForm.Get("ndc")
	lcNdc, lcErrorString = CheckNDC(lcNdc, lcErrorString)
	lcName := acRequest.PostForm.Get("name")
	lcErrorString = CheckString(lcName, lcErrorString)
	lcForm := acRequest.PostForm.Get("form")
	lcItem := acRequest.PostForm.Get("item_num")
	lcPkgSize := acRequest.PostForm.Get("pkgSize")

	if lcErrorString != "" {
		ExecuteTemplate(acWriter, "newDrug.html", lcErrorString)
		return
	}
	mainUtils.UpdateDrug(lcPkgSize, lcForm, lcItem, lcName, lcNdc)
	GetCloseHandler(acWriter, acRequest)
	return
}
