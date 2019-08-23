/**
File: getPrescriptionHandler
Description: Gets new prescription page
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
Function: GetPrescriptionHandler
Description: Execute the prescription template
*/
func GetPrescriptionHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.LogError(err.Error())
	}
	utils.ExecuteTemplate(acWriter, "prescription.html", nil)
}
