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

var templates = Must(ParseFiles("web/templates/audit.html",
	"web/templates/closeWindow.html", "web/templates/database.html",
	"web/templates/newDrug.html", "web/templates/prescription.html",
	"web/templates/purchase.html"))

/**
Function: ExecuteTemplate
Description: Executes a html template
@param w The http writer
@param tmpl The template name
@param data The data for the template to use
*/
func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
