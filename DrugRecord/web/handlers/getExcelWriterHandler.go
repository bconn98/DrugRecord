/**
File: getExcelWriterHandler
Description: Gets new excel writer page
@author Bryan Conn
@date 11/17/2019
*/
package handlers

import (
	"log"
	"net/http"

	"../../mainUtils"
	"github.com/sqweek/dialog"
)

/**
Function: GetExcelWriterHandler
Description: Executes the excel writer
*/
func GetExcelWriterHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	lcFileName, err := dialog.File().Filter("Excel Workbook (*.xlsx)", "xlsx").Title("Export to XLSX").Save()
	if err != nil {
		log.Fatal(err)
	}
	mainUtils.ExcelWriter(lcFileName)

	GetCloseHandler(acWriter, acRequest)
}
