/**
File: postAuditHandler
Description: Sends the audit information
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/bconn98/DrugRecord/mainUtils"
	"github.com/bconn98/DrugRecord/web/utils"
)

/**
Function: PostAuditHandler
Description: Sends the audit information to add it to the database and executes the
database template to refresh the page
*/
func PostAuditHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.Log(err.Error(), mainUtils.ERROR)
	}

	var lcErrorString string
	lcNdc := acRequest.PostForm.Get("ndc")
	lcNdc, lcErrorString = utils.CheckNDC(lcNdc, lcErrorString)
	lcPharmacist := acRequest.PostForm.Get("pharmacist")
	lcPharmacist = strings.ToUpper(lcPharmacist)
	lcAuditDate := acRequest.PostForm.Get("AuditDate")
	lcAuditMonth, lcAuditDay, lcAuditYear := utils.ParseDate(lcAuditDate)
	lcErrorString, lcAuditYear = utils.CheckDate(lcAuditMonth, lcAuditDay, lcAuditYear, lcErrorString)
	lnQty := acRequest.PostForm.Get("qty")

	lrQty, err := strconv.ParseFloat(lnQty, 64)
	if err != nil {
		mainUtils.Log(err.Error(), mainUtils.ERROR)
	}

	if lcErrorString != "" {
		utils.ExecuteTemplate(acWriter, "audit.html", lcErrorString)
		return
	}
	lbCheck, _ := mainUtils.AddAudit(mainUtils.MakeAudit(lcNdc, lcPharmacist, lrQty, lcAuditYear, lcAuditMonth,
		lcAuditDay))

	if !lbCheck {
		utils.ExecuteTemplate(acWriter, "audit.html", "Audit already logged!")
		return
	}

	GetCloseHandler(acWriter, acRequest)
}
