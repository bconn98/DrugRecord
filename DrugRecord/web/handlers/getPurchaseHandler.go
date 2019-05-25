/**
File: getPurchaseHandler
Description: Gets new purchase page
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
Function: GetPurchaseHandler
Description: Executes the purchase template
*/
func GetPurchaseHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	utils.ExecuteTemplate(w, "purchase.html", nil)
}
