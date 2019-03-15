/**
File: postNewDrugHandler
Description: Sends the new drug information
@author Bryan Conn
@date 1/4/2019
 */
package handlers

import (
	"net/http"
	"../../mainUtils"
	"strconv"
)

/**
Function: PostNewDrugHelper
Description: Sends the purchase information to be added to the database
and executes the database template to refresh
*/
func PostNewDrugHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ndc := r.PostForm.Get("ndc")
	name := r.PostForm.Get("name")
	form := r.PostForm.Get("form")
	itemS := r.PostForm.Get("itemnum")
	itemNum, _ := strconv.Atoi(itemS)
	pkgSize := r.PostForm.Get("pkgsize")
	mainUtils.UpdateDrug(pkgSize, form, itemNum, name, ndc)
	GetCloseHandler(w, r)
	return
}