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
Function: PostDatabaseHandler
Description: Sends the information matching the entered NDC to be executed
in the database template
*/
func PostDatabaseHandler(acWriter ResponseWriter, acRequest *Request) {
	err := acRequest.ParseForm()
	if err != nil {
		LogError(err.Error())
	}
	var lcErrorString string
	lcNdc := acRequest.PostForm.Get("ndc")
	lcNdc, lcErrorString = CheckNDC(lcNdc, lcErrorString)

	if lcErrorString != "" {
		ExecuteTemplate(acWriter, "database.html", nil)
		return
	}
	lcName, lcNdc, lcForm, lcItemNum, lcSize, lcDate, lnQty, lasOrders := FindNDC(lcNdc)

	if lcName == "" && lcNdc == "" && lcForm == "" && lcItemNum == ""  {
		ExecuteTemplate(acWriter, "database.html", nil)
		return
	}

	lcDateString := lcDate.Month().String() + " " + strconv.Itoa(lcDate.Day()) + " " + strconv.Itoa(lcDate.Year())
	lsData := data{lcName, lcNdc, lcForm, lcSize, lcDateString,
		lcItemNum, lnQty, lasOrders}
	ExecuteTemplate(acWriter, "database.html", lsData)
	return
}
