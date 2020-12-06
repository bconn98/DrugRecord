/**
File: getPrescriptionHandler
Description: Gets new prescription page
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"net/http"

	"github.com/jimlawless/whereami"

	"github.com/bconn98/DrugRecord/utils"
	"github.com/bconn98/DrugRecord/web/webUtils"
)

/**
Function: GetPrescriptionHandler
Description: Execute the prescription template
*/
func GetPrescriptionHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}
	webUtils.ExecuteTemplate(acWriter, "prescription.html", nil)
}
