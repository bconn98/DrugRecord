/**
File: postPurchaseHandler
Description: Sends the purchase information
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/bconn98/DrugRecord/mainUtils"
	"github.com/bconn98/DrugRecord/web/utils"
)

/**
Function: PostPurchaseHelper
Description: Sends the purchase information to be added to the database
and executes the database template to refresh
*/
func PostPurchaseHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.LogError(err.Error())
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
	lrQty, err := strconv.ParseFloat(acRequest.PostForm.Get("qty"), 64)
	if err != nil {
		mainUtils.LogError(err.Error())
	}

	lcActual := acRequest.PostForm.Get("realCount")
	if lcActual == "" {
		lcActual = "-1000" // Set a default value that should never be seen
	}
	lrActual, err := strconv.ParseFloat(lcActual, 64)
	if err != nil {
		mainUtils.LogError(err.Error())
	}

	if lcErrorString != "" {
		utils.ExecuteTemplate(acWriter, "purchase.html", lcErrorString)
		return
	}

	// Checks if the drug exists yet
	lbCheck := mainUtils.NewCheck(lcNdc)
	purchase := mainUtils.MakePurchase(lcNdc, lcPharmacist, lcInvoice, lcYear, lcMonth,
		lcDay, lrQty, lrActual)

	// If the drug does exist
	if lbCheck {
		lbLogged, _ := mainUtils.AddPurchase(purchase)

		if !lbLogged {
			utils.ExecuteTemplate(acWriter, "purchase.html", "Purchase already logged!")
			return
		}

		GetCloseHandler(acWriter, acRequest)
		return
	} else {
		mainUtils.AddDrug(lcNdc, lcMonth, lcDay, lcYear)
		_, id := mainUtils.AddPurchase(purchase)
		utils.ExecuteTemplate(acWriter, "newDrug.html", mainUtils.NewDrug{Ndc: lcNdc, Id: id})
	}
}
