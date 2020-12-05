/**
File: postRegisterHandler
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
Function: PostRegisterHandler
Description: Sends the new users information to be validated and redirects differently
depending on that validity.
*/
func PostRegisterHandler(acWriter http.ResponseWriter, acRequest *http.Request) {
	err := acRequest.ParseForm()
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}

	lcUsername := acRequest.PostForm.Get("uName")
	lcPassword := acRequest.PostForm.Get("password")

	leValidation := utils.MakeUser(lcUsername, lcPassword)
	switch leValidation {
	case utils.UE:
		webUtils.ExecuteTemplate(acWriter, "register.html", "Username field cannot be left empty!")
		return
	case utils.US:
		webUtils.ExecuteTemplate(acWriter, "register.html", "Username field cannot contain spaces!")
		return
	case utils.PE:
		webUtils.ExecuteTemplate(acWriter, "register.html", "Password field cannot be left empty!")
		return
	case utils.PS:
		webUtils.ExecuteTemplate(acWriter, "register.html", "Password field cannot contain spaces!")
		return
	case utils.TN:
		webUtils.ExecuteTemplate(acWriter, "register.html", "This username is already in use, please try again!")
		return
	case utils.GOOD:
		http.Redirect(acWriter, acRequest, "/login", http.StatusFound)
	}
}
