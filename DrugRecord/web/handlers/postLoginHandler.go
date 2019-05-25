/**
File: postLoginHandler
Description: Sends the audit information
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	. "../../mainUtils"
	. "../utils"
	"log"
	. "net/http"
)

/**
Function: PostLoginHandler
Description: Sends the login information for validation, redirects depending on the outcome
*/
func PostLoginHandler(w ResponseWriter, r *Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	username := r.PostForm.Get("uName")
	user := FindUser(username)
	test := User{}
	if user == test {
		ExecuteTemplate(w, "login.html", "Unknown Username!")
		return
	}
	password := r.PostForm.Get("password")
	if !CheckPassword(user, password) {
		ExecuteTemplate(w, "login.html", "Password doesn't match our records!")
		return
	}
	SetGood()
	Redirect(w, r, "database", 302)
	return
}
