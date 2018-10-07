package handlers

import (
	"net/http"
	"../../mainUtils"
)

func PostPrescriptionHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	udc := r.PostForm.Get("udc")
	pharmacist := r.PostForm.Get("pharmacist")
	//_ := r.PostForm.Get("script")
	month := r.PostForm.Get("month")
	day := r.PostForm.Get("day")
	year := r.PostForm.Get("year")
	qty := r.PostForm.Get("qty")
	mainUtils.AddPrescription(udc, pharmacist, month, day, year, qty)
	GetDatabaseHandler(w, r)
}