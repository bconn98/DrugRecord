package handlers

import (
	. "net/http"
	. "../utils"
)

func PostLoginHandler(w ResponseWriter, r *Request) {
	r.ParseForm()
	username := r.PostForm.Get("uName")
	user := FindUser(username)
	test := User{}
	if user == test {
		ExecuteTemplate(w, "login.html", "Unknown Username!")
		return
	}
	password := r.PostForm.Get("password")
	if CheckPassword(user, password) {
		ExecuteTemplate(w, "home.html", r)
		return
	}
	ExecuteTemplate(w, "login.html", "Password doesn't match our records!")
	return
}