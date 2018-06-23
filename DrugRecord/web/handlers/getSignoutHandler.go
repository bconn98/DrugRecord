package handlers

import (
	"net/http"
	"../utils"
)

func GetSignoutHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "home.html", nil)
}