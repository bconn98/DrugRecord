/**
File: getHomeHandler
Description: Gets new home page
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	"net/http"
	"../utils"
)

/**
Function: GetHomeHandler
Description: Executes the home template
*/
func GetHomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "home.html", nil)
}