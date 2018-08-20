package handlers

import (
	"net/http"
	"../utils"
)

func GetPurchaseHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "purchase.html", nil)
}