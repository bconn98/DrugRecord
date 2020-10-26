/**
File: postExcelWriterHandler
Description: Sends the excel writer information
@author Bryan Conn
@date 11/17/2019
*/
package handlers

import (
	"net/http"

	"github.com/jimlawless/whereami"

	"github.com/bconn98/DrugRecord/utils"
)

/**
Function: PostExcelWriterHandler
Description: Sends the audit information to add it to the database and executes the
database template to refresh the page
*/
func PostExcelWriterHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}

	lcFileName := acRequest.PostForm.Get("fileName")

	utils.ExcelWriter(lcFileName)

	GetCloseHandler(acWriter, acRequest)
}
