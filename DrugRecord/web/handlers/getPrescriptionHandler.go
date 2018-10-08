/**
File: getPrescriptionHandler
Description: Gets new prescription page
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	"net/http"
	"../utils"
)

/**
Function: GetPrescriptionHandler
Description: Execute the prescription template
 */
func GetPrescriptionHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "prescription.html", nil)
}