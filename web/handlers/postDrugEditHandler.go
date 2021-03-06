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

	"github.com/jimlawless/whereami"

	"github.com/bconn98/DrugRecord/utils"
	. "github.com/bconn98/DrugRecord/web/webUtils"
)

/**
Function: PostDeleteHandler
Description: Sends the delete information to find the order to delete
*/
func PostDrugEditHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
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
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}

	utils.UpdateDrug(lcPkgSize, lcForm, lcItem, lcName, lcNdc, lrQty, lcNdc)
	GetCloseHandler(acWriter, acRequest)
	return
}
