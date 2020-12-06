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

	lcErrStr := "N/A"
	lcTemplateName := "register.html"
	leValidation := utils.MakeUser(lcUsername, lcPassword)
	switch leValidation {
	case utils.UE:
		lcErrStr = "Username field cannot be left empty!"
		break
	case utils.US:
		lcErrStr = "Username field cannot contain spaces!"
		break
	case utils.PE:
		lcErrStr = "Password field cannot be left empty!"
		break
	case utils.PS:
		lcErrStr = "Password field cannot contain spaces!"
		break
	case utils.PIV:
		lcErrStr = "Password length must be between 10 and 72 characters"
		break
	case utils.TN:
		lcErrStr = "This username is already in use, please try again!"
		break
	case utils.GOOD:
		http.Redirect(acWriter, acRequest, "/login", http.StatusFound)
	}

	if lcErrStr != "N/A" {
		webUtils.ExecuteTemplate(acWriter, lcTemplateName, lcErrStr)
	}
}
