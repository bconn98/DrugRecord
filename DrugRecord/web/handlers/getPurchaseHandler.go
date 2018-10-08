/**
File: getPurchaseHandler
Description: Gets new purchase page
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	"net/http"
	"../utils"
)

/**
Function: GetPurchaseHandler
Description: Executes the purchase template
 */
func GetPurchaseHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "purchase.html", nil)
}