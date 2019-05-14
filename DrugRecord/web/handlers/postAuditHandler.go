/**
File: postAuditHandler
Description: Sends the audit information
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"../../mainUtils"
	"../utils"
	"log"
	"net/http"
)

/**
Function: PostAuditHandler
Description: Sends the audit information to add it to the database and executes the
database template to refresh the page
*/
func PostAuditHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	var str string
	ndc := r.PostForm.Get("ndc")
	ndc, str = utils.CheckNDC(ndc, str)
	pharmacist := r.PostForm.Get("pharmacist")
	amonth := r.PostForm.Get("amonth")
	aday := r.PostForm.Get("aday")
	ayear := r.PostForm.Get("ayear")
	str = utils.CheckDate(amonth, aday, ayear, str)
	qty := r.PostForm.Get("qty")
	actual := r.PostForm.Get("realCount")
	str = utils.CheckQty(actual, str)
	if str != "" {
		utils.ExecuteTemplate(w, "audit.html", str)
		return
	}
	check := mainUtils.AddAudit(ndc, pharmacist, amonth, aday, ayear, qty, actual)

	if !check {
		utils.ExecuteTemplate(w, "audit.html", "Audit already logged!")
		return
	}

	GetCloseHandler(w, r)
}
