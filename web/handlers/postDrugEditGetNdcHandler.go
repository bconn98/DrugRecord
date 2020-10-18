/**
File: postDeleteHandler
Description: Sends the audit information
@author Bryan Conn
@date 6/2/2019
*/
package handlers

import (
	"net/http"

	"github.com/bconn98/DrugRecord/utils"
	"github.com/bconn98/DrugRecord/web/webUtils"
)

/**
Function: PostDeleteHandler
Description: Sends the delete information to find the order to delete
*/
func PostDrugEditGetNdcHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR)
	}

	var lcErrorString string
	lcNdc := acRequest.PostForm.Get("ndc")
	lcNdc, _ = webUtils.CheckNDC(lcNdc, lcErrorString)

	drug := utils.GetDrug(lcNdc)

	webUtils.ExecuteTemplate(acWriter, "editDrug.html", drug)
	return
}
