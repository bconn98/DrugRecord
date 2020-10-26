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

	"github.com/jimlawless/whereami"

	"github.com/bconn98/DrugRecord/utils"
	"github.com/bconn98/DrugRecord/web/webUtils"
)

/**
Function: PostPurchaseHelper
Description: Sends the purchase information to be added to the database
and executes the database template to refresh
*/
func PostPurchaseHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}

	var lcErrorString string
	lcNdc := acRequest.PostForm.Get("ndc")
	lcNdc, lcErrorString = webUtils.CheckNDC(lcNdc, lcErrorString)
	lcPharmacist := acRequest.PostForm.Get("pharmacist")
	lcPharmacist = strings.ToUpper(lcPharmacist)
	lcInvoice := acRequest.PostForm.Get("invoice")
	lcPurchaseDate := acRequest.PostForm.Get("PurchaseDate")
	lcMonth, lcDay, lcYear := webUtils.ParseDate(lcPurchaseDate)
	lcErrorString, lcYear = webUtils.CheckDate(lcMonth, lcDay, lcYear, lcErrorString)
	lrQty, err := strconv.ParseFloat(acRequest.PostForm.Get("qty"), 64)
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}

	lcActual := acRequest.PostForm.Get("realCount")
	if lcActual == "" {
		lcActual = "-1000" // Set a default value that should never be seen
	}
	lrActual, err := strconv.ParseFloat(lcActual, 64)
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}

	if lcErrorString != "" {
		webUtils.ExecuteTemplate(acWriter, "purchase.html", lcErrorString)
		return
	}

	// Checks if the drug exists yet
	lbCheck := utils.NewCheck(lcNdc)
	purchase := utils.MakePurchase(lcNdc, lcPharmacist, lcInvoice, lcYear, lcMonth,
		lcDay, lrQty, lrActual)

	// If the drug does exist
	if lbCheck {
		lbLogged, _ := utils.AddPurchase(purchase)

		if !lbLogged {
			webUtils.ExecuteTemplate(acWriter, "purchase.html", "Purchase already logged!")
			return
		}

		GetCloseHandler(acWriter, acRequest)
		return
	} else {
		utils.AddDrug(lcNdc, lcMonth, lcDay, lcYear)
		_, id := utils.AddPurchase(purchase)
		webUtils.ExecuteTemplate(acWriter, "newDrug.html", utils.NewDrug{Ndc: lcNdc, Id: id})
	}
}
