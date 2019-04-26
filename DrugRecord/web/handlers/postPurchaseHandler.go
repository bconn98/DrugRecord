/**
File: postPurchaseHandler
Description: Sends the purchase information
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
	ndc, str = utils.CheckNDC(ndc, str)
	pharmacist := r.PostForm.Get("pharmacist")
	order := r.PostForm.Get("order")
	str = utils.CheckNum(order, str)
	month := r.PostForm.Get("month")
	day := r.PostForm.Get("day")
	year := r.PostForm.Get("year")
	str = utils.CheckDate(month, day, year, str)
	qty := r.PostForm.Get("qty")
	str = utils.CheckQty(qty, str)
	actual := r.PostForm.Get("realCount")
	str = utils.CheckQty(actual, str)
	if str != "" {
		utils.ExecuteTemplate(w, "purchase.html", str)
		return
	}
	// Checks if the drug exists yet
	check := mainUtils.NewCheck(ndc)
	// If the drug does exist
	if check {
		mainUtils.AddPurchase(ndc, pharmacist, month, day, year, qty, order, actual)
		GetCloseHandler(w, r)
		return
	} else {
		mainUtils.AddDrug(ndc, month, day, year)
		utils.ExecuteTemplate(w, "newDrug.html", nil)
		mainUtils.AddPurchase(ndc, pharmacist, month, day, year, qty, order, actual)
	}
}