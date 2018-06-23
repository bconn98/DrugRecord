package handlers

import (
	"net/http"
	"../utils"
)

func GetLoginHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}