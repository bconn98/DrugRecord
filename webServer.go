/**
File: webServer
Description: Runs a database
@author Bryan Conn
@date 10/7/2018
*/
package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gorilla/mux"
	"github.com/jimlawless/whereami"
	"gopkg.in/go-ini/ini.v1"

	"github.com/bconn98/DrugRecord/utils"
	"github.com/bconn98/DrugRecord/web/handlers"
)

/**
Function: main
Description: Creates a new web server at port 8080 and connects all of the handler functions
*/
func main() {
	var err error

	utils.Initial = true

	lcIniFile, err := ini.Load("configs/configuration.ini")
	if err != nil {
		log.Fatal("Failed to open configuration file")
	}

	lcLogLevel := lcIniFile.Section("Logging").Key("log_level").String()
	utils.McHost = lcIniFile.Section("PostgreSQL").Key("host").String()
	utils.McPort = lcIniFile.Section("PostgreSQL").Key("port").String()
	utils.McDatabase = lcIniFile.Section("PostgreSQL").Key("database").String()
	utils.McUsername = lcIniFile.Section("PostgreSQL").Key("username").String()
	utils.McPassword = lcIniFile.Section("PostgreSQL").Key("password").String()

	switch lcLogLevel {
	case "DEBUG":
		utils.GbLogLevel = utils.DEBUG
		break
	case "SQL":
		utils.GbLogLevel = utils.SQL
		break
	case "INFO":
		utils.GbLogLevel = utils.INFO
		break
	case "WARNING":
		utils.GbLogLevel = utils.WARNING
		break
	case "ERROR":
		utils.GbLogLevel = utils.ERROR
		break
	}

	// If the file doesn't exist, create it, or append to the file
	utils.Log("Starting Program", utils.INFO, whereami.WhereAmI())

	utils.McConnStr = "postgres://" + utils.McUsername + ":" + utils.
		McPassword + "@" + utils.McHost + "/" + utils.McDatabase + "?sslmode=disable"
	utils.McDb, err = sql.Open("postgres", utils.McConnStr)
	if err != nil {
		utils.Log("Failed to connect to database", utils.FATAL, whereami.WhereAmI())
		log.Fatal("Failed to connect to database: " + err.Error())
	}

	if err = utils.McDb.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	lcGoScheduler := gocron.NewScheduler(time.UTC)
	_, err = lcGoScheduler.Every(1).Day().At("04:00").Do(utils.Backup)
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}
	lcGoScheduler.StartAsync()

	defer func() {
		if err = utils.GpcFile.Close(); err != nil {
			log.Fatal(err)
		}
		if err = utils.McDb.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	utils.McRouter = mux.NewRouter()
	utils.McRouter.HandleFunc("/home", handlers.GetHomeHandler).Methods("GET")
	utils.McRouter.HandleFunc("/SignOut", handlers.GetSignOutHandler).Methods("GET")
	utils.McRouter.HandleFunc("/closeWindow", handlers.GetCloseHandler).Methods("GET")
	utils.McRouter.HandleFunc("/writeExcel", handlers.GetExcelWriterHandler).Methods("GET")
	utils.McRouter.HandleFunc("/delete", handlers.PostDeleteHandler).Methods("POST")

	utils.McRouter.HandleFunc("/edit/{id:[0-9]+}", handlers.GetEditHandler).Methods("GET")
	utils.McRouter.HandleFunc("/editQty", handlers.PostEditQtyHandler).Methods("POST")

	utils.McRouter.HandleFunc("/editDrug", handlers.GetDrugEditHandler).Methods("GET")
	utils.McRouter.HandleFunc("/editDrug", handlers.PostDrugEditHandler).Methods("POST")

	utils.McRouter.HandleFunc("/editDrugGetNdc", handlers.GetDrugEditGetNdcHandler).Methods("GET")
	utils.McRouter.HandleFunc("/editDrugGetNdc", handlers.PostDrugEditGetNdcHandler).Methods("POST")

	utils.McRouter.HandleFunc("/delete/{id:[0-9]+}", handlers.GetDeleteHandler).Methods("GET")

	utils.McRouter.HandleFunc("/newDrug", handlers.GetNewDrugHandler).Methods("GET")
	utils.McRouter.HandleFunc("/newDrug", handlers.PostNewDrugHandler).Methods("POST")

	utils.McRouter.HandleFunc("/", handlers.GetDatabaseNdcHandler).Methods("GET")
	utils.McRouter.HandleFunc("/databaseDrug", handlers.GetDatabaseNdcHandler).Methods("GET")
	utils.McRouter.HandleFunc("/databaseDrug", handlers.PostDatabaseNdcHandler).Methods("POST")

	utils.McRouter.HandleFunc("/databaseDrug/{ndc:[0-9]{5}-[0-9]{4}-[0-9]{2}}",
		handlers.GetDatabaseNdcClickHandler).Methods("GET")

	utils.McRouter.HandleFunc("/databaseName", handlers.GetDatabaseNameHandler).Methods("GET")
	utils.McRouter.HandleFunc("/databaseName", handlers.PostDatabaseNameHandler).Methods("POST")

	utils.McRouter.HandleFunc("/audit", handlers.GetAuditHandler).Methods("GET")
	utils.McRouter.HandleFunc("/audit", handlers.PostAuditHandler).Methods("POST")

	utils.McRouter.HandleFunc("/purchase", handlers.GetPurchaseHandler).Methods("GET")
	utils.McRouter.HandleFunc("/purchase", handlers.PostPurchaseHandler).Methods("POST")

	utils.McRouter.HandleFunc("/prescription", handlers.GetPrescriptionHandler).Methods("GET")
	utils.McRouter.HandleFunc("/prescription", handlers.PostPrescriptionHandler).Methods("POST")

	// mainUtils.McRouter.HandleFunc("/login", handlers.GetLoginHandler).Methods("GET")
	// mainUtils.McRouter.HandleFunc("/login", handlers.PostLoginHandler).Methods("POST")

	// mainUtils.McRouter.HandleFunc("/register", handlers.GetRegisterHandler).Methods("GET")
	// mainUtils.McRouter.HandleFunc("/register", handlers.PostRegisterHandler).Methods("POST")

	http.Handle("/web/assets/", http.StripPrefix("/web/assets", http.FileServer(http.Dir("./web/assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", utils.McRouter)
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		utils.Log("Failed to listen on port 80", utils.ERROR, whereami.WhereAmI())
		log.Fatal(err)
	}
}
