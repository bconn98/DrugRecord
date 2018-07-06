package handlers

import (
	"net/http"
	"../utils"
	"../../mainUtils"
)

func GetDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	users := mainUtils.GetUsers()
	utils.ExecuteTemplate(w,"database.html", users)
}