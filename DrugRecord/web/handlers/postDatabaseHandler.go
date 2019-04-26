/**
File: postDatabaseHandler
Description: Sends the database information
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	. "../../mainUtils"
	. "../utils"
	. "net/http"
	"strconv"
)

type data struct {
	Name, Ndc, Form, Size, Date string
	ItemNum string
	Qty int
	Orders []Order
}

/**
Function: PostDatabaseHandler
Description: Sends the information matching the entered NDC to be executed
in the database template
*/
func PostDatabaseHandler(w ResponseWriter, r *Request) {
	r.ParseForm()
	var str string
	ndc := r.PostForm.Get("ndc")
	ndc, str = CheckNDC( ndc, str )

	if str != "" {
		ExecuteTemplate(w, "database.html", nil)
		return
	}
	name, ndc, form, itemNum, size, date, qty, orders := FindNDC(ndc)
	dateS := date.Month().String() + " " + strconv.Itoa(date.Day()) + " " + strconv.Itoa(date.Year())
	dataD := data{name, ndc, form, size, dateS, itemNum, qty, orders}
	ExecuteTemplate(w,"database.html", dataD)
	return
}