package handlers

import (
	"net/http"
	"../utils"
)

func GetAuditHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "audit.html", nil)
}