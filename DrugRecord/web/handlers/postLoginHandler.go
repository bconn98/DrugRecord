/**
File: postLoginHandler
Description: Sends the audit information
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"log"
	. "net/http"

	. "../../mainUtils"
	. "../utils"
)

/**
Function: PostLoginHandler
Description: Sends the login information for validation, redirects depending on the outcome
*/
func PostLoginHandler(acWriter ResponseWriter, acRequest *Request) {
	err := acRequest.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	lcUsername := acRequest.PostForm.Get("uName")
	lsUser := FindUser(lcUsername)
	lsTestUser := User{}
	if lsUser == lsTestUser {
		ExecuteTemplate(acWriter, "login.html", "Unknown Username!")
		return
	}
	lcPassword := acRequest.PostForm.Get("password")
	if !CheckPassword(lsUser, lcPassword) {
		ExecuteTemplate(acWriter, "login.html", "Password doesn't match our records!")
		return
	}
	SetSignedIn()
	Redirect(acWriter, acRequest, "database", 302)
	return
}
