/**
File: postLoginHandler
Description: Sends the audit information
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"net/http"

	"github.com/bconn98/DrugRecord/mainUtils"
	"github.com/bconn98/DrugRecord/web/utils"
)

/**
Function: PostLoginHandler
Description: Sends the login information for validation, redirects depending on the outcome
*/
func PostLoginHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.Log(err.Error(), mainUtils.ERROR)
	}
	lcUsername := acRequest.PostForm.Get("uName")
	lsUser := mainUtils.FindUser(lcUsername)
	lsTestUser := mainUtils.User{}
	if lsUser == lsTestUser {
		utils.ExecuteTemplate(acWriter, "login.html", "Unknown Username!")
		return
	}
	lcPassword := acRequest.PostForm.Get("password")
	if !mainUtils.CheckPassword(lsUser, lcPassword) {
		utils.ExecuteTemplate(acWriter, "login.html", "Password doesn't match our records!")
		return
	}
	SetSignedIn()
	http.Redirect(acWriter, acRequest, "database", 302)
	return
}
