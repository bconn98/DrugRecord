package handlers

import (
	"net/http"
	"../utils"
	"../../mainUtils"
)

func GetDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	users := mainUtils.FindUDC("")
	utils.ExecuteTemplate(w,"database.html", users)
}