/**
File: postRegisterHandler
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
Function: PostRegisterHandler
Description: Sends the new users information to be validated and redirects differently
depending on that validity.
*/
func PostRegisterHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		mainUtils.Log(err.Error(), mainUtils.ERROR)
	}

	lcUsername := acRequest.PostForm.Get("uName")
	lcPassword := acRequest.PostForm.Get("password")

	leValidation := mainUtils.MakeUser(lcUsername, lcPassword)
	switch leValidation {
	case mainUtils.UE:
		utils.ExecuteTemplate(acWriter, "register.html", "Username field cannot be left empty!")
		return
	case mainUtils.US:
		utils.ExecuteTemplate(acWriter, "register.html", "Username field cannot contain spaces!")
		return
	case mainUtils.PE:
		utils.ExecuteTemplate(acWriter, "register.html", "Password field cannot be left empty!")
		return
	case mainUtils.PS:
		utils.ExecuteTemplate(acWriter, "register.html", "Password field cannot contain spaces!")
		return
	case mainUtils.TN:
		utils.ExecuteTemplate(acWriter, "register.html", "This username is already in use, please try again!")
		return
	case mainUtils.GOOD:
		http.Redirect(acWriter, acRequest, "/login", http.StatusFound)
	}
}
