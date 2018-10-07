package handlers

import (
	"net/http"
	"../../mainUtils"
)

func PostAuditHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	udc := r.PostForm.Get("udc")
	pharmacist := r.PostForm.Get("pharmacist")
	amonth := r.PostForm.Get("amonth")
	aday := r.PostForm.Get("aday")
	ayear := r.PostForm.Get("ayear")
	qty := r.PostForm.Get("qty")
	mainUtils.AddAudit(udc, pharmacist, amonth, aday, ayear, qty)
	GetDatabaseHandler(w, r)
}