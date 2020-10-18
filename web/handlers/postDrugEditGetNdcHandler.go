/**
File: postDeleteHandler
Description: Sends the audit information
@author Bryan Conn
@date 6/2/2019
*/
package handlers

import (
	"net/http"

	"github.com/bconn98/DrugRecord/mainUtils"
	. "github.com/bconn98/DrugRecord/web/utils"
)

/**
Function: PostDeleteHandler
Description: Sends the delete information to find the order to delete
*/
func PostDrugEditGetNdcHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.Log(err.Error(), mainUtils.ERROR)
	}

	var lcErrorString string
	lcNdc := acRequest.PostForm.Get("ndc")
	lcNdc, _ = CheckNDC(lcNdc, lcErrorString)

	drug := mainUtils.GetDrug(lcNdc)

	ExecuteTemplate(acWriter, "editDrug.html", drug)
	return
}
