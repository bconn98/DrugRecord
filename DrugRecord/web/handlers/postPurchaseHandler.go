/**
File: postPurchaseHandler
Description: Sends the audit information
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	"net/http"
	"../../mainUtils"
	"../utils"
)

/**
Function: PostPurchaseHelper
Description: Sends the purchase information to be added to the database
and executes the database template to refresh
*/
func PostPurchaseHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	udc := r.PostForm.Get("udc")
	pharmacist := r.PostForm.Get("pharmacist")
	month := r.PostForm.Get("month")
	day := r.PostForm.Get("day")
	year := r.PostForm.Get("year")
	qty := r.PostForm.Get("qty")
	mainUtils.AddPurchase(udc, pharmacist, month, day, year, qty)
	utils.ExecuteTemplate(w, "closeWindow.html", nil)
}