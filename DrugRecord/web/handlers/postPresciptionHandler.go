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
)

/**
Function: PostPrescriptionHandler
Description: Sends the prescription information to be added to the database and executes the
database template to refresh
*/
func PostPrescriptionHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	udc := r.PostForm.Get("udc")
	pharmacist := r.PostForm.Get("pharmacist")
	//_ := r.PostForm.Get("script")
	month := r.PostForm.Get("month")
	day := r.PostForm.Get("day")
	year := r.PostForm.Get("year")
	qty := r.PostForm.Get("qty")
	mainUtils.AddPrescription(udc, pharmacist, month, day, year, qty)
	GetDatabaseHandler(w, r)
}