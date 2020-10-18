/**
File: postPrescriptionHandler
Description: Sends the prescription information
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
Function: PostPrescriptionHandler
Description: Sends the prescription information to be added to the database and executes the
database template to refresh
*/
func PostPrescriptionHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.Log(err.Error(), mainUtils.ERROR)
	}

	var lcErrorString string
	lcNdc := acRequest.PostForm.Get("ndc")
	lcNdc, lcErrorString = utils.CheckNDC(lcNdc, lcErrorString)
	lcPharmacist := acRequest.PostForm.Get("pharmacist")
	lcPharmacist = strings.ToUpper(lcPharmacist)
	lcScript := acRequest.PostForm.Get("script")
	lcOrderDate := acRequest.PostForm.Get("OrderDate")
	lcMonth, lcDay, lcYear := utils.ParseDate(lcOrderDate)
	lcErrorString, lcYear = utils.CheckDate(lcMonth, lcDay, lcYear, lcErrorString)
	lnQty := acRequest.PostForm.Get("qty")
	lnActual := acRequest.PostForm.Get("realCount")

	lrQty, err := strconv.ParseFloat(lnQty, 64)
	if err != nil {
		mainUtils.Log(err.Error(), mainUtils.ERROR)
	}

	lrActual, err := strconv.ParseFloat(lnActual, 64)
	if err != nil {
		mainUtils.Log(err.Error(), mainUtils.ERROR)
	}

	if lnActual == "" {
		lrActual = -1000
	}

	if lcErrorString != "" {
		utils.ExecuteTemplate(acWriter, "prescription.html", lcErrorString)
		return
	}

	lbCheck := mainUtils.NewCheck(lcNdc)
	prescription := mainUtils.MakePrescription(lcNdc, lcPharmacist, lcScript, lrQty, lcYear, lcMonth, lcDay, lrActual)

	// If the drug does exist
	if lbCheck {

		lbLogged, _ := mainUtils.AddPrescription(prescription)

		if !lbLogged {
			utils.ExecuteTemplate(acWriter, "prescription.html", "Prescription already logged!")
			return
		}

		GetCloseHandler(acWriter, acRequest)
		return
	} else {
		mainUtils.AddDrug(lcNdc, lcMonth, lcDay, lcYear)
		_, id := mainUtils.AddPrescription(prescription)
		utils.ExecuteTemplate(acWriter, "newDrug.html", mainUtils.NewDrug{Ndc: lcNdc, Id: id})
	}
}
