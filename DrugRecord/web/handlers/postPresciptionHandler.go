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
func PostPrescriptionHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	var str string
	ndc := r.PostForm.Get("ndc")
	ndc, str = utils.CheckNDC(ndc, str)
	pharmacist := r.PostForm.Get("pharmacist")
	script := r.PostForm.Get("script")
	month := r.PostForm.Get("month")
	day := r.PostForm.Get("day")
	year := r.PostForm.Get("year")
	str = utils.CheckDate(month, day, year, str)
	qty := r.PostForm.Get("qty")
	actual := r.PostForm.Get("realCount")
	str = utils.CheckQty(actual, str)
	if str != "" {
		utils.ExecuteTemplate(w, "prescription.html", str)
		return
	}
	check := mainUtils.AddPrescription(ndc, pharmacist, month, day, year, qty, script, actual)

	if !check {
		utils.ExecuteTemplate(w, "prescription.html", "Prescription already logged!")
		return
	}

	GetCloseHandler(w, r)
}
