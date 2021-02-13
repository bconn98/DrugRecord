package utils

import (
	"database/sql"

	"github.com/gorilla/mux"
)

var McRouter *mux.Router

var McHost string

var McPort string

var McDatabase string

var McUsername string

var McPassword string

// Database password needs to be changed when released
var McConnStr string
var McDb *sql.DB
