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

	"github.com/jimlawless/whereami"

	"github.com/bconn98/DrugRecord/utils"
	"github.com/bconn98/DrugRecord/web/webUtils"
)

/**
Function: PostPrescriptionHandler
Description: Sends the prescription information to be added to the database and executes the
database template to refresh
*/
func PostPrescriptionHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}

	var lcErrorString string
	lcNdc := acRequest.PostForm.Get("ndc")
	lcNdc, lcErrorString = webUtils.CheckNDC(lcNdc, lcErrorString)
	lcPharmacist := acRequest.PostForm.Get("pharmacist")
	lcPharmacist = strings.ToUpper(lcPharmacist)
	lcScript := acRequest.PostForm.Get("script")
	lcOrderDate := acRequest.PostForm.Get("OrderDate")
	lcMonth, lcDay, lcYear := webUtils.ParseDate(lcOrderDate)
	lcErrorString, lcYear = webUtils.CheckDate(lcMonth, lcDay, lcYear, lcErrorString)
	lnQty := acRequest.PostForm.Get("qty")
	lnActual := acRequest.PostForm.Get("realCount")

	lrQty, err := strconv.ParseFloat(lnQty, 64)
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}

	lrActual, err := strconv.ParseFloat(lnActual, 64)
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}

	if lnActual == "" {
		lrActual = -1000
	}

	if lcErrorString != "" {
		webUtils.ExecuteTemplate(acWriter, "prescription.html", lcErrorString)
		return
	}

	lbCheck := utils.NewCheck(lcNdc)
	prescription := utils.MakePrescription(lcNdc, lcPharmacist, lcScript, lrQty, lcYear, lcMonth, lcDay, lrActual)

	// If the drug does exist
	if lbCheck {

		lbLogged, _ := utils.AddPrescription(prescription)

		if !lbLogged {
			webUtils.ExecuteTemplate(acWriter, "prescription.html", "Prescription already logged!")
			return
		}

		GetCloseHandler(acWriter, acRequest)
		return
	} else {
		utils.AddDrug(lcNdc, lcMonth, lcDay, lcYear)
		_, id := utils.AddPrescription(prescription)
		webUtils.ExecuteTemplate(acWriter, "newDrug.html", utils.NewDrug{Ndc: lcNdc, Id: id})
	}
}
