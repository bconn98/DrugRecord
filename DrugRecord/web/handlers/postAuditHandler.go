/**
File: postAuditHandler
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
Function: PostAuditHandler
Description: Sends the audit information to add it to the database and executes the
database template to refresh the page
*/
func PostAuditHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ndc := r.PostForm.Get("ndc")
	pharmacist := r.PostForm.Get("pharmacist")
	amonth := r.PostForm.Get("amonth")
	aday := r.PostForm.Get("aday")
	ayear := r.PostForm.Get("ayear")
	qty := r.PostForm.Get("qty")
	mainUtils.AddAudit(ndc, pharmacist, amonth, aday, ayear, qty)
	GetCloseHandler(w, r)
}