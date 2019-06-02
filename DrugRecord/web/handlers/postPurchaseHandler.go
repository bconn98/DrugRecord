/**
File: postPurchaseHandler
Description: Sends the purchase information
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"log"
	"net/http"
	"strings"

	"../../mainUtils"
	"../utils"
)

/**
Function: PostPurchaseHelper
Description: Sends the purchase information to be added to the database
and executes the database template to refresh
*/
func PostPurchaseHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var lcErrorString string
	lcNdc := acRequest.PostForm.Get("ndc")
	lcNdc, lcErrorString = utils.CheckNDC(lcNdc, lcErrorString)
	lcPharmacist := acRequest.PostForm.Get("pharmacist")
	lcPharmacist = strings.ToUpper(lcPharmacist)
	lcInvoice := acRequest.PostForm.Get("invoice")
	lcPurchaseDate := acRequest.PostForm.Get("PurchaseDate")
	lcMonth, lcDay, lcYear := utils.ParseDate(lcPurchaseDate)
	lcErrorString, lcYear = utils.CheckDate(lcMonth, lcDay, lcYear, lcErrorString)
	lnQty := acRequest.PostForm.Get("qty")
	lnActual := acRequest.PostForm.Get("realCount")
	if lcErrorString != "" {
		utils.ExecuteTemplate(acWriter, "purchase.html", lcErrorString)
		return
	}

	// Checks if the drug exists yet
	lbCheck := mainUtils.NewCheck(lcNdc)
	// If the drug does exist
	if lbCheck {
		lbLogged := mainUtils.AddPurchase(lcNdc, lcPharmacist, lcMonth, lcDay, lcYear, lnQty, lcInvoice, lnActual)

		if !lbLogged {
			utils.ExecuteTemplate(acWriter, "purchase.html", "Purchase already logged!")
			return
		}

		GetCloseHandler(acWriter, acRequest)
		return
	} else {
		mainUtils.AddDrug(lcNdc, lcMonth, lcDay, lcYear)
		utils.ExecuteTemplate(acWriter, "newDrug.html", nil)
		mainUtils.AddPurchase(lcNdc, lcPharmacist, lcMonth, lcDay, lcYear, lnQty, lcInvoice, lnActual)
	}
}
