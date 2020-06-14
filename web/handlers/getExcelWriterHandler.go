/**
File: getExcelWriterHandler
Description: Gets new excel writer page
@author Bryan Conn
@date 11/17/2019
*/
package handlers

import (
	"net/http"

	"github.com/sqweek/dialog"

	"github.com/bconn98/DrugRecord/mainUtils"
)

/**
Function: GetExcelWriterHandler
Description: Executes the excel writer
*/
func GetExcelWriterHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	lcFileName, err := dialog.File().Filter("Excel Workbook (*.xlsx)", "xlsx").Title("Export to XLSX").Save()
	if err != nil {
		mainUtils.LogError(err.Error())
	}
	mainUtils.ExcelWriter(lcFileName)

	GetCloseHandler(acWriter, acRequest)
}
