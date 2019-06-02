/**
File: postPrescriptionHandler
Description: Sends the prescription information
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"log"
	"net/http"

	"../../mainUtils"
	"../utils"
)

/**
Function: PostPrescriptionHandler
Description: Sends the prescription information to be added to the database and executes the
database template to refresh
*/
func PostPrescriptionHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var lcErrorString string
	lcNdc := acRequest.PostForm.Get("ndc")
	lcNdc, lcErrorString = utils.CheckNDC(lcNdc, lcErrorString)
	lcPharmacist := acRequest.PostForm.Get("pharmacist")
	lcScript := acRequest.PostForm.Get("script")
	lcOrderDate := acRequest.PostForm.Get("OrderDate")
	lcMonth, lcDay, lcYear := utils.ParseDate(lcOrderDate)
	lcErrorString, lcYear = utils.CheckDate(lcMonth, lcDay, lcYear, lcErrorString)
	lnQty := acRequest.PostForm.Get("qty")
	lnActual := acRequest.PostForm.Get("realCount")
	lcErrorString = utils.CheckQty(lnActual, lcErrorString)
	if lcErrorString != "" {
		utils.ExecuteTemplate(acWriter, "prescription.html", lcErrorString)
		return
	}

	lbCheck := mainUtils.NewCheck(lcNdc)
	// If the drug does exist
	if lbCheck {

		lbLogged := mainUtils.AddPrescription(lcNdc, lcPharmacist, lcMonth, lcDay, lcYear, lnQty, lcScript, lnActual)

		if !lbLogged {
			utils.ExecuteTemplate(acWriter, "prescription.html", "Prescription already logged!")
			return
		}

		GetCloseHandler(acWriter, acRequest)
		return
	} else {
		mainUtils.AddDrug(lcNdc, lcMonth, lcDay, lcYear)
		utils.ExecuteTemplate(acWriter, "newDrug.html", nil)
		mainUtils.AddPrescription(lcNdc, lcPharmacist, lcMonth, lcDay, lcYear, lnQty, lcScript, lnActual)
	}
}
