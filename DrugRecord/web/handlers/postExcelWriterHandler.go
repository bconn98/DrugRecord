/**
File: postExcelWriterHandler
Description: Sends the excel writer information
@author Bryan Conn
@date 11/17/2019
*/
package handlers

import (
	"net/http"

	"../../mainUtils"
)

/**
Function: PostExcelWriterHandler
Description: Sends the audit information to add it to the database and executes the
database template to refresh the page
*/
func PostExcelWriterHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.LogError(err.Error())
	}

	lcFileName := acRequest.PostForm.Get("fileName")

	mainUtils.ExcelWriter(lcFileName)

	GetCloseHandler(acWriter, acRequest)
}