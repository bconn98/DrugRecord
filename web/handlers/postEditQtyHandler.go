/**
File: postEditQtyHandler
Description: Sends the audit information
@author Bryan Conn
@date 6/1/2019
*/
package handlers

import (
	"net/http"

	"github.com/bconn98/DrugRecord/mainUtils"
)

/**
Function: PostEditQtyHandler
Description: Sends the edit information to find the order to edit also edits the script
*/
func PostEditQtyHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.Log(err.Error(), mainUtils.ERROR)
	}

	lnId := acRequest.PostForm.Get("id")
	lcScript := acRequest.PostForm.Get("script")
	lnQty := acRequest.PostForm.Get("quantity")

	mainUtils.UpdateOrder(lnId, lcScript, lnQty)
	GetCloseHandler(acWriter, acRequest)
}
