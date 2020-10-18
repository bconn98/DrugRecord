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

	"github.com/bconn98/DrugRecord/utils"
	"github.com/bconn98/DrugRecord/web/webUtils"
)

/**
Function: PostAuditHandler
Description: Sends the audit information to add it to the database and executes the
database template to refresh the page
*/
func PostAuditHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR)
	}

	var lcErrorString string
	lcNdc := acRequest.PostForm.Get("ndc")
	lcNdc, lcErrorString = webUtils.CheckNDC(lcNdc, lcErrorString)
	lcPharmacist := acRequest.PostForm.Get("pharmacist")
	lcPharmacist = strings.ToUpper(lcPharmacist)
	lcAuditDate := acRequest.PostForm.Get("AuditDate")
	lcAuditMonth, lcAuditDay, lcAuditYear := webUtils.ParseDate(lcAuditDate)
	lcErrorString, lcAuditYear = webUtils.CheckDate(lcAuditMonth, lcAuditDay, lcAuditYear, lcErrorString)
	lnQty := acRequest.PostForm.Get("qty")

	lrQty, err := strconv.ParseFloat(lnQty, 64)
	if err != nil {
		utils.Log(err.Error(), utils.ERROR)
	}

	if lcErrorString != "" {
		webUtils.ExecuteTemplate(acWriter, "audit.html", lcErrorString)
		return
	}
	lbCheck, _ := utils.AddAudit(utils.MakeAudit(lcNdc, lcPharmacist, lrQty, lcAuditYear, lcAuditMonth,
		lcAuditDay))

	if !lbCheck {
		webUtils.ExecuteTemplate(acWriter, "audit.html", "Audit already logged!")
		return
	}

	GetCloseHandler(acWriter, acRequest)
}
