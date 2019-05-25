/**
File: getPrescriptionHandler
Description: Gets new prescription page
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"../utils"
	"log"
	"net/http"
)

/**
Function: GetPrescriptionHandler
Description: Execute the prescription template
*/
func GetPrescriptionHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	utils.ExecuteTemplate(w, "prescription.html", nil)
}
