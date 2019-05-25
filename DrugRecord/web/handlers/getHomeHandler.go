/**
File: getHomeHandler
Description: Gets new home page
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
Function: GetHomeHandler
Description: Executes the home template
*/
func GetHomeHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	utils.ExecuteTemplate(w, "home.html", nil)
}
