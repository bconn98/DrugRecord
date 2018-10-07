package handlers

import (
	"net/http"
	"../utils"
	"../../mainUtils"
)

func GetDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	users := mainUtils.FindNDC("12345-6789-12")
	utils.ExecuteTemplate(w,"database.html", users)
}