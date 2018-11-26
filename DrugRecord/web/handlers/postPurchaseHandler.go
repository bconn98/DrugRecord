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
)

/**
Function: PostPurchaseHelper
Description: Sends the purchase information to be added to the database
and executes the database template to refresh
*/
func PostPurchaseHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ndc := r.PostForm.Get("ndc")
	pharmacist := r.PostForm.Get("pharmacist")
	month := r.PostForm.Get("month")
	day := r.PostForm.Get("day")
	year := r.PostForm.Get("year")
	qty := r.PostForm.Get("qty")
	mainUtils.AddPurchase(ndc, pharmacist, month, day, year, qty)
	GetCloseHandler(w, r)
}