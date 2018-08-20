package handlers

import (
	. "net/http"
	. "../utils"
	. "../../mainUtils"
)

func PostDatabaseHandler(w ResponseWriter, r *Request) {
	r.ParseForm()
	udc := r.PostForm.Get("udc")
	orders := FindUDC(udc)
	ExecuteTemplate(w,"database.html", orders)
	return
}