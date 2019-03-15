/**
File: getNewDrugHandler
Description: Gets new drug page
@author Bryan Conn
@date 1/4/2019
 */
package handlers

import (
"net/http"
"../utils"
)

/**
Function: GetNewDrugHandler
Description: Executes the new drug template
*/
func GetNewDrugHandler(w http.ResponseWriter, r *http.Request) {
utils.ExecuteTemplate(w, "newDrug.html", nil)
}