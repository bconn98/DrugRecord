/**
File: postEditQtyHandler
Description: Sends the audit information
@author Bryan Conn
@date 6/1/2019
*/
package handlers

import (
	"log"
	"net/http"

	"../../mainUtils"
)

/**
Function: PostEditQtyHandler
Description: Sends the edit information to find the order to edit also edits the script
*/
func PostEditQtyHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	lnId := acRequest.PostForm.Get("id")
	lcScript := acRequest.PostForm.Get("script")
	lnQty := acRequest.PostForm.Get("quantity")

	mainUtils.UpdateOrder(lnId, lcScript, lnQty)
	GetCloseHandler(acWriter, acRequest)
}
