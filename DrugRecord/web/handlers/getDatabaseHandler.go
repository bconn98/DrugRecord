package handlers

import (
	"net/http"
	"../utils"
	"../../SQLDB"
)

func GetDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	users := SQLDB.GetUsers()
	utils.ExecuteTemplate(w,"database.html", users)
}