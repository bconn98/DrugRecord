package handlers

import (
	. "net/http"
	. "../utils"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("SU93R-S3CR3T"))

func PostRegisterHandler(w ResponseWriter, r *Request) {
	r.ParseForm()
	username := r.PostForm.Get("uName")
	sess, _ := store.Get(r, "session")
	sess.Values["username"] = username
	sess.Save(r, w)
	test := User{}
	if FindUser(username) != test {
		ExecuteTemplate(w, "register.html", "This username is already in use, please try again!")
	}
	password := r.PostForm.Get("password")
	MakeUser(username, password)
	Redirect(w, r, "/login", StatusFound)
}
