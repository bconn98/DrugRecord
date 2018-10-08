/**
File: getRegisterHandler
Description: Gets new audit page
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	"net/http"
	"../utils"
)

/**
Function: GetRegisterHandler
Description: Execute the register template
*/
func GetRegisterHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}