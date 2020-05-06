/**
File: postDatabaseHandler
Description: Sends the database information
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	. "net/http"
	"strconv"

	. "../../mainUtils"
	. "../utils"
)

type data struct {
	Name, Ndc, Form, Size, Date string
	ItemNum                     string
	Qty                         float64
	Orders                      []Order
}

/**
Function: PostDatabaseNdcHandler
Description: Sends the information matching the entered NDC to be executed
in the database template
*/
func PostDatabaseNdcHandler(acWriter ResponseWriter, acRequest *Request) {
	err := acRequest.ParseForm()
	if err != nil {
		LogError(err.Error())
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
		lcInput, lcErrorString = CheckNDC(lcInput, lcErrorString)

		if lcErrorString != "" {
			ExecuteTemplate(acWriter, "databaseDrug.html", nil)
			return
		}
		lcName, lcInput, lcForm, lcItemNum, lcSize, lcDate, lnQty, lasOrders := FindNDC(lcInput)

		if lcName == "" && lcInput == "" && lcForm == "" && lcItemNum == "" {
			ExecuteTemplate(acWriter, "databaseDrug.html", nil)
			return
		}

		lcDateString := lcDate.Month().String() + " " + strconv.Itoa(lcDate.Day()) + " " + strconv.Itoa(lcDate.Year())
		lsData := data{lcName, lcInput, lcForm, lcSize, lcDateString,
			lcItemNum, lnQty, lasOrders}
		ExecuteTemplate(acWriter, "databaseDrug.html", lsData)
		return
	} else {
		lasDrugs := GetDrugs(lcInput)
		lsData := dataName{lasDrugs}
		ExecuteTemplate(acWriter, "databaseName.html", lsData)
	}
}
