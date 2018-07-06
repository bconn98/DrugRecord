package handlers

import (
	. "net/http"
	. "../utils"
	. "../../mainUtils"
	"github.com/gorilla/sessions"
	"../../SQLDB"
	"strings"
)

var store = sessions.NewCookieStore([]byte("SU93R-S3CR3T"))

func PostRegisterHandler(w ResponseWriter, r *Request) {
	r.ParseForm()
	username := r.PostForm.Get("uName")
	password := r.PostForm.Get("password")
	if username == "" {
		ExecuteTemplate(w, "register.html", "Username field cannot be left empty!")
		return
	} else if strings.Contains(username, " ") {
		ExecuteTemplate(w, "register.html", "Username field cannot contain spaces!")
		return
	} else if password == "" {
		ExecuteTemplate(w, "register.html", "Password field cannot be left empty!")
		return
	} else if strings.Contains(password, " ") {
		ExecuteTemplate(w, "register.html", "Password field cannot contain spaces!")
		return
	}
	sess, _ := store.Get(r, "session")
	sess.Values["username"] = username
	sess.Save(r, w)
	test := User{}
	if FindUser(username) != test {
		ExecuteTemplate(w, "register.html", "This username is already in use, please try again!")
		return
	}
	MakeUser(username, password)
	SQLDB.AddUser(username, password)
	Redirect(w, r, "/login", StatusFound)
}
