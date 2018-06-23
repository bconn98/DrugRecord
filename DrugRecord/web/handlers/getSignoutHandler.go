package handlers

import (
	"net/http"
	"../utils"
)

func GetSignoutHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "learning.html", nil)
}