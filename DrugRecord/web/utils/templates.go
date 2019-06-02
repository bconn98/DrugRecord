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
	"web/templates/closeWindow.html", "web/templates/database.html",
	"web/templates/newDrug.html", "web/templates/prescription.html",
	"web/templates/purchase.html", "web/templates/edit.html",
	"web/templates/editQty.html", "web/templates/delete.html",
    "web/templates/deleteSure.html"))

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
