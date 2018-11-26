/**
File: getSignoutHandler
Description: Gets new signout page
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	"net/http"
	"../utils"
)

/**
Function: GetSignoutHandler
Description: Executes the signout template
*/
func GetSignoutHandler(w http.ResponseWriter, r *http.Request) {
	SetBad()
	utils.ExecuteTemplate(w, "home.html", nil)
}