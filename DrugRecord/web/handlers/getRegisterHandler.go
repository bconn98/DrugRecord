/**
File: getRegisterHandler
Description: Gets new audit page
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"../utils"
	"log"
	"net/http"
)

/**
Function: GetRegisterHandler
Description: Execute the register template
*/
func GetRegisterHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	utils.ExecuteTemplate(w, "register.html", nil)
}
