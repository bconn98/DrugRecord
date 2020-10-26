/**
File: postNewDrugHandler
Description: Sends the new drug information
@author Bryan Conn
@date 1/4/2019
*/
package handlers

import (
	"net/http"

	"github.com/jimlawless/whereami"

	"github.com/bconn98/DrugRecord/utils"
	"github.com/bconn98/DrugRecord/web/webUtils"
)

/**
Function: PostNewDrugHelper
Description: Sends the purchase information to be added to the database
and executes the database template to refresh
*/
func PostNewDrugHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}

	var lcErrorString string
	lcId := acRequest.PostForm.Get("id")
	lcOldNdc := acRequest.PostForm.Get("oldNdc")
	lcNdc := acRequest.PostForm.Get("ndc")
	lcNdc, lcErrorString = webUtils.CheckNDC(lcNdc, lcErrorString)
	lcName := acRequest.PostForm.Get("name")
	lcErrorString = webUtils.CheckString(lcName, lcErrorString)
	lcForm := acRequest.PostForm.Get("form")
	lcItem := acRequest.PostForm.Get("item_num")
	lcPkgSize := acRequest.PostForm.Get("pkgSize")

	if lcErrorString != "" {

		if utils.NewCheck(lcNdc) {
			utils.UpdateOrderNdc(lcId, lcNdc)
			GetCloseHandler(acWriter, acRequest)
			return
		}

		webUtils.ExecuteTemplate(acWriter, "newDrug.html", utils.NewDrug{Error: lcErrorString, Ndc: lcNdc})
		return
	}
	utils.UpdateDrug(lcPkgSize, lcForm, lcItem, lcName, lcNdc, 0, lcOldNdc)
	utils.UpdateOrderNdc(lcId, lcNdc)
	GetCloseHandler(acWriter, acRequest)
	return
}
