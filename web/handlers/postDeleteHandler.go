/**
File: postDeleteHandler
Description: Sends the audit information
@author Bryan Conn
@date 6/2/2019
*/
package handlers

import (
	"net/http"
	"strconv"

	"github.com/bconn98/DrugRecord/mainUtils"
)

/**
Function: PostDeleteHandler
Description: Sends the delete information to find the order to delete
*/
func PostDeleteHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.LogError(err.Error())
	}

	lcId := acRequest.PostForm.Get("id")
	lnId, err := strconv.ParseInt(lcId, 10, 64)

	mainUtils.DeleteOrder(lnId)
	GetCloseHandler(acWriter, acRequest)
}
