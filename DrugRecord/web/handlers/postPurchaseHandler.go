/**
File: postPurchaseHandler
Description: Sends the purchase information
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"../../mainUtils"
	"../utils"
	"log"
	"net/http"
)

/**
Function: PostPurchaseHelper
Description: Sends the purchase information to be added to the database
and executes the database template to refresh
*/
func PostPurchaseHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	var str string
	ndc := r.PostForm.Get("ndc")
	ndc, str = utils.CheckNDC(ndc, str)
	pharmacist := r.PostForm.Get("pharmacist")
	invoice := r.PostForm.Get("invoice")
	month := r.PostForm.Get("month")
	day := r.PostForm.Get("day")
	year := r.PostForm.Get("year")
	str = utils.CheckDate(month, day, year, str)
	qty := r.PostForm.Get("qty")
	actual := r.PostForm.Get("realCount")
	str = utils.CheckQty(actual, str)
	if str != "" {
		utils.ExecuteTemplate(w, "purchase.html", str)
		return
	}

	// Checks if the drug exists yet
	check := mainUtils.NewCheck(ndc)
	// If the drug does exist
	if check {
		logged := mainUtils.AddPurchase(ndc, pharmacist, month, day, year, qty, invoice, actual)

		if !logged {
			utils.ExecuteTemplate(w, "purchase.html", "Purchase already logged!")
			return
		}

		GetCloseHandler(w, r)
		return
	} else {
		mainUtils.AddDrug(ndc, month, day, year)
		utils.ExecuteTemplate(w, "newDrug.html", nil)
		mainUtils.AddPurchase(ndc, pharmacist, month, day, year, qty, invoice, actual)
	}
}
