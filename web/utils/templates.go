/**
File: templates
Description: Executes templates and implements the range for html
@author Bryan Conn
@date: 10/7/2018
*/
package utils

import (
	. "html/template"
	"net/http"
)

var mcTemplates = Must(ParseFiles("web/templates/audit.html",
	"web/templates/closeWindow.html", "web/templates/databaseDrug.html",
	"web/templates/newDrug.html", "web/templates/prescription.html",
	"web/templates/purchase.html", "web/templates/editQty.html",
<<<<<<< HEAD:web/utils/templates.go
	"web/templates/delete.html", "web/templates/writeExcel.html",
=======
	"web/templates/deleteSure.html", "web/templates/writeExcel.html",
>>>>>>> master:DrugRecord/web/utils/templates.go
	"web/templates/editDrug.html", "web/templates/editDrugGetNdc.html",
	"web/templates/databaseName.html"))

/**
Function: ExecuteTemplate
Description: Executes a html template
@param acWriter The http writer
@param acTemplate The template name
@param aiData The data for the template to use
*/
func ExecuteTemplate(acWriter http.ResponseWriter, acTemplate string, aiData interface{}) {
	err := mcTemplates.ExecuteTemplate(acWriter, acTemplate, aiData)
	if err != nil {
		http.Error(acWriter, err.Error(), http.StatusInternalServerError)
	}
}
