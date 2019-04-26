/**
File: postPrescriptionHandler
Description: Sends the prescription information
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	"net/http"
	"../../mainUtils"
	"../utils"
)

/**
Function: PostPrescriptionHandler
Description: Sends the prescription information to be added to the database and executes the
database template to refresh
*/
func PostPrescriptionHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var str string
	ndc := r.PostForm.Get("ndc")
	ndc, str = utils.CheckNDC(ndc, str)
	pharmacist := r.PostForm.Get("pharmacist")
	script := r.PostForm.Get("script")
	str = utils.CheckNum(script, str)
	month := r.PostForm.Get("month")
	day := r.PostForm.Get("day")
	year := r.PostForm.Get("year")
	str = utils.CheckDate(month, day, year, str)
	qty := r.PostForm.Get("qty")
	str = utils.CheckQty(qty, str)
	actual := r.PostForm.Get("realCount")
	str = utils.CheckQty(actual, str)
	if str != "" {
		utils.ExecuteTemplate(w, "prescription.html", str)
		return
	}
	mainUtils.AddPrescription(ndc, pharmacist, month, day, year, qty, script, actual)
	GetCloseHandler(w, r)
}