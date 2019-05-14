/**
File: postRegisterHandler
Description: Sends the audit information
@author Bryan Conn
@date 10/7/18
 */
package handlers

import (
	"log"
	. "net/http"
	. "../utils"
	. "../../mainUtils"
)

/**
Function: PostRegisterHandler
Description: Sends the new users information to be validated and redirects differently
depending on that validity.
*/
func PostRegisterHandler(w ResponseWriter, r *Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	username := r.PostForm.Get("uName")
	password := r.PostForm.Get("password")

	validation := MakeUser(username, password)
	switch validation {
		case UE:
			ExecuteTemplate(w, "register.html", "Username field cannot be left empty!")
			return
		case US:
			ExecuteTemplate(w, "register.html", "Username field cannot contain spaces!")
			return
		case PE:
			ExecuteTemplate(w, "register.html", "Password field cannot be left empty!")
			return
		case PS:
			ExecuteTemplate(w, "register.html", "Password field cannot contain spaces!")
			return
		case TN:
			ExecuteTemplate(w, "register.html", "This username is already in use, please try again!")
			return
		case GOOD:
			Redirect(w, r, "/login", StatusFound)
	}
}
