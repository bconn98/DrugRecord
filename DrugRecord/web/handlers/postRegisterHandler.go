/**
File: postRegisterHandler
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
Function: PostRegisterHandler
Description: Sends the new users information to be validated and redirects differently
depending on that validity.
*/
func PostRegisterHandler(acWriter ResponseWriter, acRequest *Request) {
	err := acRequest.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	lcUsername := acRequest.PostForm.Get("uName")
	lcPassword := acRequest.PostForm.Get("password")

	leValidation := MakeUser(lcUsername, lcPassword)
	switch leValidation {
	case UE:
		ExecuteTemplate(acWriter, "register.html", "Username field cannot be left empty!")
		return
	case US:
		ExecuteTemplate(acWriter, "register.html", "Username field cannot contain spaces!")
		return
	case PE:
		ExecuteTemplate(acWriter, "register.html", "Password field cannot be left empty!")
		return
	case PS:
		ExecuteTemplate(acWriter, "register.html", "Password field cannot contain spaces!")
		return
	case TN:
		ExecuteTemplate(acWriter, "register.html", "This username is already in use, please try again!")
		return
	case GOOD:
		Redirect(acWriter, acRequest, "/login", StatusFound)
	}
}
