/**
File: getLoginHandler
Description: Gets new login page
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	"net/http"
	"../utils"
)
/**
Function: GetLoginHandler
Description: Executes the login template
*/
func GetLoginHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}