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

	"github.com/jimlawless/whereami"

	"github.com/bconn98/DrugRecord/utils"
)

/**
Function: PostDeleteHandler
Description: Sends the delete information to find the order to delete
*/
func PostDeleteHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}

	lcId := acRequest.PostForm.Get("id")
	lnId, err := strconv.ParseInt(lcId, 10, 64)

	utils.DeleteOrder(lnId)
	GetCloseHandler(acWriter, acRequest)
}
