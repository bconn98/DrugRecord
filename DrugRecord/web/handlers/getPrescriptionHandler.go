package handlers

import (
	"net/http"
	"../utils"
)

func GetPrescriptionHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "prescription.html", nil)
}