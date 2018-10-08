/**
File: postLoginHandler
Description: Sends the audit information
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	. "net/http"
	. "../../mainUtils"
	. "../utils"
)

/**
Function: PostLoginHandler
Description: Sends the login information for validation, redirects depending on the outcome
*/
func PostLoginHandler(w ResponseWriter, r *Request) {
	r.ParseForm()
	username := r.PostForm.Get("uName")
	user := FindUser(username)
	test := User{}
	if user == test {
		ExecuteTemplate(w, "login.html", "Unknown Username!")
		return
	}
	password := r.PostForm.Get("password")
	if!CheckPassword(user, password) {
		ExecuteTemplate(w, "login.html", "Password doesn't match our records!")
		return
	}
	users := GetUsers()
	ExecuteTemplate(w,"database.html", users)
	return
}