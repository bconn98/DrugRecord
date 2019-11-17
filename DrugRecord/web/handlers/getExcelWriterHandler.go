/**
File: getExcelWriterHandler
Description: Gets new excel writer page
@author Bryan Conn
@date 11/17/2019
*/
package handlers

import (
	"net/http"

	"../../mainUtils"
	"../utils"
)

/**
Function: GetExcelWriterHandler
Description: Executes the audit template
*/
func GetExcelWriterHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.LogError(err.Error())
	}
	utils.ExecuteTemplate(acWriter, "writeExcel.html", nil)
}
