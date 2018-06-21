package handlers

import (
	"net/http"
	"../utils"
)
func GetRegisterHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}