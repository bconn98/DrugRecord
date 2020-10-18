/**
File: postDatabaseHandler
Description: Sends the database information
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"net/http"
	"strconv"

	"github.com/bconn98/DrugRecord/utils"
	"github.com/bconn98/DrugRecord/web/webUtils"
)

type data struct {
	Name, Ndc, Form, Size, Date string
	ItemNum                     string
	Qty                         float64
	Orders                      []utils.Order
}

/**
Function: PostDatabaseNdcHandler
Description: Sends the information matching the entered NDC to be executed
in the database template
*/
func PostDatabaseNdcHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR)
	}
	var lcErrorString string
	lcInput := acRequest.PostForm.Get("search")

	var ndc = true
	for i := 0; i < len(lcInput); i++ {
		if _, err := strconv.ParseInt(string(lcInput[i]), 10, 64); err != nil && lcInput[i] != '-' {
			ndc = false
		}
	}

	if ndc {
		lcInput, lcErrorString = webUtils.CheckNDC(lcInput, lcErrorString)

		if lcErrorString != "" {
			webUtils.ExecuteTemplate(acWriter, "databaseDrug.html", nil)
			return
		}
		lcName, lcInput, lcForm, lcItemNum, lcSize, lcDate, lnQty, lasOrders := utils.FindNDC(lcInput)

		if lcName == "" && lcInput == "" && lcForm == "" && lcItemNum == "" {
			webUtils.ExecuteTemplate(acWriter, "databaseDrug.html", nil)
			return
		}

		lcDateString := lcDate.Month().String() + " " + strconv.Itoa(lcDate.Day()) + " " + strconv.Itoa(lcDate.Year())
		lsData := data{lcName, lcInput, lcForm, lcSize, lcDateString,
			lcItemNum, lnQty, lasOrders}

		webUtils.ExecuteTemplate(acWriter, "databaseDrug.html", lsData)
		return
	} else {
		lasDrugs := utils.GetDrugs(lcInput)
		lsData := dataName{lasDrugs}
		webUtils.ExecuteTemplate(acWriter, "databaseName.html", lsData)
	}
}
