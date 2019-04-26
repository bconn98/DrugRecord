/**
File: postNewDrugHandler
Description: Sends the new drug information
@author Bryan Conn
@date 1/4/2019
 */
package handlers

import (
	"../../mainUtils"
	"net/http"
	. "../utils"
)

/**
Function: PostNewDrugHelper
Description: Sends the purchase information to be added to the database
and executes the database template to refresh
*/
func PostNewDrugHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var str string
	ndc := r.PostForm.Get("ndc")
	ndc, str = CheckNDC( ndc, str)
	name := r.PostForm.Get("name")
	str = CheckString(name, str)
	form := r.PostForm.Get("form")
	itemS := r.PostForm.Get("itemnum")
	pkgSize := r.PostForm.Get("pkgsize")

	if str != "" {
		ExecuteTemplate(w, "newDrug.html", str)
		return
	}
	mainUtils.UpdateDrug(pkgSize, form, itemS, name, ndc)
	GetCloseHandler(w, r)
	return
}