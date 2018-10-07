package handlers

import (
	"net/http"
	"../../mainUtils"
)

func PostPurchaseHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	udc := r.PostForm.Get("udc")
	pharmacist := r.PostForm.Get("pharmacist")
	month := r.PostForm.Get("month")
	day := r.PostForm.Get("day")
	year := r.PostForm.Get("year")
	qty := r.PostForm.Get("qty")
	mainUtils.AddPurchase(udc, pharmacist, month, day, year, qty)
	GetDatabaseHandler(w, r)
}