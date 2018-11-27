/**
File: postPurchaseHandler
Description: Sends the audit information
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	"../utils"
	"net/http"
	"../../mainUtils"
)

/**
Function: PostPurchaseHelper
Description: Sends the purchase information to be added to the database
and executes the database template to refresh
*/
func PostPurchaseHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var str string
	ndc := r.PostForm.Get("ndc")
	str = utils.CheckNDC(ndc, str)
	pharmacist := r.PostForm.Get("pharmacist")
	str = utils.CheckPharm(pharmacist, str)
	order := r.PostForm.Get("order")
	str = utils.CheckNum(order, str)
	month := r.PostForm.Get("month")
	day := r.PostForm.Get("day")
	year := r.PostForm.Get("year")
	str = utils.CheckDate(month, day, year, str)
	qty := r.PostForm.Get("qty")
	str = utils.CheckQty(qty, str)
	if str != "" {
		utils.ExecuteTemplate(w, "purchase.html", str)
		return
	}
	mainUtils.AddPurchase(ndc, pharmacist, month, day, year, qty)
	GetCloseHandler(w, r)
}