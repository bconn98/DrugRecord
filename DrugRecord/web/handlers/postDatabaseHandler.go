package handlers

import (
	. "net/http"
	. "../utils"
	. "../../mainUtils"
)

func PostDatabaseHandler(w ResponseWriter, r *Request) {
	r.ParseForm()
	ndc := r.PostForm.Get("ndc")
	orders := FindNDC(ndc)
	ExecuteTemplate(w,"database.html", orders)
	return
}