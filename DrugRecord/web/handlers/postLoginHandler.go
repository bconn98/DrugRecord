package handlers


import (
	"net/http"
	"../utils"
)

func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "learning.html", nil)
}