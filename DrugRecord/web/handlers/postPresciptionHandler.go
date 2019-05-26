/**
File: postPrescriptionHandler
Description: Sends the prescription information
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
	lcMonth := acRequest.PostForm.Get("month")
	lcDay := acRequest.PostForm.Get("day")
	lcYear := acRequest.PostForm.Get("year")
	lcErrorString = utils.CheckDate(lcMonth, lcDay, lcYear, lcErrorString)
	lnQty := acRequest.PostForm.Get("qty")
	lnActual := acRequest.PostForm.Get("realCount")
	lcErrorString = utils.CheckQty(lnActual, lcErrorString)
	if lcErrorString != "" {
		utils.ExecuteTemplate(acWriter, "prescription.html", lcErrorString)
		return
	}
	check := mainUtils.AddPrescription(lcNdc, lcPharmacist, lcMonth, lcDay, lcYear, lnQty, lcScript, lnActual)

	if !check {
		utils.ExecuteTemplate(acWriter, "prescription.html", "Prescription already logged!")
		return
	}

	GetCloseHandler(acWriter, acRequest)
}
