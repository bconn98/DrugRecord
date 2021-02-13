/**
File: postLoginHandler
Description: Sends the audit information
@author Bryan Conn
@date 10/7/18
*/
package handlers

import (
	"net/http"

	"github.com/jimlawless/whereami"

	utils "github.com/bconn98/DrugRecord/utils"
	webUtils "github.com/bconn98/DrugRecord/web/webUtils"
)

/**
Function: PostLoginHandler
Description: Sends the login information for validation, redirects depending on the outcome
*/
func PostLoginHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}
	lcUsername := acRequest.PostForm.Get("uName")
	lsUser := utils.FindUser(lcUsername)
	lsTestUser := utils.User{}
	if lsUser == lsTestUser {
		webUtils.ExecuteTemplate(acWriter, "login.html", "Unknown Username!")
		return
	}
	lcPassword := acRequest.PostForm.Get("password")
	if !utils.CheckPassword(lsUser, lcPassword) {
		webUtils.ExecuteTemplate(acWriter, "login.html", "Password doesn't match our records!")
		return
	}
	SetSignedIn()
	http.Redirect(acWriter, acRequest, "database", 302)
	return
}
