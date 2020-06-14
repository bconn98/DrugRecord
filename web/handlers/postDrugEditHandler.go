/**
File: postDeleteHandler
Description: Sends the audit information
@author Bryan Conn
@date 6/2/2019
*/
package handlers

import (
	"net/http"
	"strconv"

	"github.com/bconn98/DrugRecord/mainUtils"
	. "github.com/bconn98/DrugRecord/web/utils"
)

/**
Function: PostDeleteHandler
Description: Sends the delete information to find the order to delete
*/
func PostDrugEditHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.LogError(err.Error())
	}

	var lcErrorString string
	lcNdc := acRequest.PostForm.Get("ndc")
	lcNdc, _ = CheckNDC(lcNdc, lcErrorString)
	lcName := acRequest.PostForm.Get("name")
	lcForm := acRequest.PostForm.Get("form")
	lcItem := acRequest.PostForm.Get("item_num")
	lcPkgSize := acRequest.PostForm.Get("size")
	lcQty := acRequest.PostForm.Get("qty")

	lrQty, err := strconv.ParseFloat(lcQty, 10)
	if err != nil {
		mainUtils.LogError(err.Error())
	}

	mainUtils.UpdateDrug(lcPkgSize, lcForm, lcItem, lcName, lcNdc, lrQty, lcNdc)
	GetCloseHandler(acWriter, acRequest)
	return
}
