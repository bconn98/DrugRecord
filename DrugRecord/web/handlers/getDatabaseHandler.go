/**
File: getDatabaseHandler
Description: Gets a new database page
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	"net/http"
	"../utils"
	"../../mainUtils"
)

/**
Function: GetDatabaseHandler
Description: Executes the database template with the output data
*/
func GetDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	output := mainUtils.FindNDC("12345-6789-12")
	utils.ExecuteTemplate(w,"database.html", output)
}