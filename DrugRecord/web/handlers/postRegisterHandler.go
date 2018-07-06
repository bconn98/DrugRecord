package handlers

import (
	. "net/http"
	. "../utils"
	. "../../mainUtils"
)

func PostRegisterHandler(w ResponseWriter, r *Request) {
	r.ParseForm()
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
