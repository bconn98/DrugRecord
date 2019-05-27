/**
File: getPurchaseHandler
Description: Gets new purchase page
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"log"
	"net/http"

	"../utils"
)

/**
Function: GetPurchaseHandler
Description: Executes the purchase template
*/
func GetPurchaseHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	utils.ExecuteTemplate(acWriter, "purchase.html", nil)
}
