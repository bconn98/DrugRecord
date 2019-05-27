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
func PostAuditHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var lcErrorString string
	lcNdc := acRequest.PostForm.Get("ndc")
	lcNdc, lcErrorString = utils.CheckNDC(lcNdc, lcErrorString)
	lcPharmacist := acRequest.PostForm.Get("pharmacist")
	lcAuditMonth := acRequest.PostForm.Get("AuditMonth")
	lcAuditDay := acRequest.PostForm.Get("AuditDay")
	lcAuditYear := acRequest.PostForm.Get("AuditYear")
	lcErrorString = utils.CheckDate(lcAuditMonth, lcAuditDay, lcAuditYear, lcErrorString)
	lnQty := acRequest.PostForm.Get("qty")
	lnActual := acRequest.PostForm.Get("realCount")
	lcErrorString = utils.CheckQty(lnActual, lcErrorString)
	if lcErrorString != "" {
		utils.ExecuteTemplate(acWriter, "audit.html", lcErrorString)
		return
	}
	lbCheck := mainUtils.AddAudit(lcNdc, lcPharmacist, lcAuditMonth, lcAuditDay, lcAuditYear, lnQty, lnActual)

	if !lbCheck {
		utils.ExecuteTemplate(acWriter, "audit.html", "Audit already logged!")
		return
	}

	GetCloseHandler(acWriter, acRequest)
}
