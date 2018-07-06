package handlers

import (
	. "net/http"
	. "../../mainUtils"
	. "../utils"
	"../../SQLDB"
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
		users := SQLDB.GetUsers()
		ExecuteTemplate(w,"database.html", users)
		return
	}
	ExecuteTemplate(w, "login.html", "Password doesn't match our records!")
	return
}