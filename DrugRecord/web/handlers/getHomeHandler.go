package handlers

import (
	"net/http"
	"../utils"
)

func GetHomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "home.html", nil)
}