/**
File: postAuditHandler
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
Function: PostAuditHandler
Description: Sends the audit information to add it to the database and executes the
database template to refresh the page
*/
func PostAuditHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var str string
	ndc := r.PostForm.Get("ndc")
	str = utils.CheckNDC(ndc, str)
	pharmacist := r.PostForm.Get("pharmacist")
	amonth := r.PostForm.Get("amonth")
	aday := r.PostForm.Get("aday")
	ayear := r.PostForm.Get("ayear")
	str = utils.CheckDate(amonth, aday, ayear, str)
	qty := r.PostForm.Get("qty")
	str = utils.CheckQty(qty, str)
	if str != "" {
		utils.ExecuteTemplate(w, "audit.html", str)
		return
	}
	mainUtils.AddAudit(ndc, pharmacist, amonth, aday, ayear, qty)
	GetCloseHandler(w, r)
}