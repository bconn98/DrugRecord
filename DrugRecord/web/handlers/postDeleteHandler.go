/**
File: postDeleteHandler
Description: Sends the audit information
@author Bryan Conn
@date 6/2/2019
*/
package handlers

import (
	"net/http"
	"strings"

	"../../mainUtils"
	"../utils"
)

/**
Function: PostDeleteHandler
Description: Sends the delete information to find the order to delete
*/
func PostDeleteHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.LogError(err.Error())
	}

	var lcErrorString string
	lcNdc := acRequest.PostForm.Get("ndc")
	lcNdc, lcErrorString = utils.CheckNDC(lcNdc, lcErrorString)
	lcPharmacist := acRequest.PostForm.Get("pharmacist")
	lcPharmacist = strings.ToUpper(lcPharmacist)
	lcScript := acRequest.PostForm.Get("script")
	lcType := acRequest.PostForm.Get("type")
	lcDate := acRequest.PostForm.Get("date")
	lcMonth, lcDay, lcYear := utils.ParseDate(lcDate)
	lcErrorString, lcYear = utils.CheckDate(lcMonth, lcDay, lcYear, lcErrorString)
	if lcErrorString != "" {
		utils.ExecuteTemplate(acWriter, "delete.html", lcErrorString)
		return
	}
	order := mainUtils.MakeOrder(lcNdc, lcPharmacist, lcScript, lcType, 0, 0, lcYear, lcMonth, lcDay,
		0) // ID doesn't matter
	lasOrders := mainUtils.GetOrder(order)

	if len(lasOrders) != 0 {
		utils.ExecuteTemplate(acWriter, "deleteSure.html", lasOrders[0])
	} else {
		GetCloseHandler(acWriter, acRequest)
	}
}
