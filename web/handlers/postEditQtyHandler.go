/**
File: postEditQtyHandler
Description: Sends the audit information
@author Bryan Conn
@date 6/1/2019
*/
package handlers

import (
	"net/http"

	"github.com/jimlawless/whereami"

	"github.com/bconn98/DrugRecord/utils"
)

/**
Function: PostEditQtyHandler
Description: Sends the edit information to find the order to edit also edits the script
*/
func PostEditQtyHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}

	lnId := acRequest.PostForm.Get("id")
	lcScript := acRequest.PostForm.Get("script")
	lnQty := acRequest.PostForm.Get("quantity")

	utils.UpdateOrder(lnId, lcScript, lnQty)
	GetCloseHandler(acWriter, acRequest)
}
