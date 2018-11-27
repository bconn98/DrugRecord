/**
File: postDatabaseHandler
Description: Sends the database information
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	. "net/http"
	. "../utils"
	. "../../mainUtils"
)

/**
Function: PostDatabaseHandler
Description: Sends the information matching the entered NDC to be executed
in the database template
*/
func PostDatabaseHandler(w ResponseWriter, r *Request) {
	r.ParseForm()
	var str string
	ndc := r.PostForm.Get("ndc")

	if CheckNDC(ndc, str) != "" {
		ExecuteTemplate(w, "database.html", nil)
		return
	}
	orders := FindNDC(ndc)
	ExecuteTemplate(w,"database.html", orders)
	return
}