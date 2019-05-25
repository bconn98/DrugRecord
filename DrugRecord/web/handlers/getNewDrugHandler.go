/**
File: getNewDrugHandler
Description: Gets new drug page
@author Bryan Conn
@date 1/4/2019
*/
package handlers

import (
	"../utils"
	"log"
	"net/http"
)

/**
Function: GetNewDrugHandler
Description: Executes the new drug template
*/
func GetNewDrugHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	utils.ExecuteTemplate(w, "newDrug.html", nil)
}
