/**
File: webServer
Description: Runs a database
@author Bryan Conn
@date 10/7/2018
*/
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
<<<<<<< HEAD
	"gopkg.in/go-ini/ini.v1"
=======
>>>>>>> master

	"github.com/bconn98/DrugRecord/mainUtils"
	"github.com/bconn98/DrugRecord/web/handlers"
)

/**
Function: main
Description: Creates a new web server at port 8080 and connects all of the handler functions
*/
func main() {
	var err error

	mainUtils.Initial = true

	lcIniFile, err := ini.Load("configs/configuration.ini")
	if err != nil {
		log.Fatal("Failed to open configuration file")
	}

	lcLogLevel := lcIniFile.Section("Logging").Key("log_level").String()

	switch lcLogLevel {
	case "DEBUG":
		mainUtils.GbLogLevel = mainUtils.DEBUG
		break
	case "SQL":
		mainUtils.GbLogLevel = mainUtils.SQL
		break
	case "INFO":
		mainUtils.GbLogLevel = mainUtils.INFO
		break
	case "WARNING":
		mainUtils.GbLogLevel = mainUtils.WARNING
		break
	case "ERROR":
		mainUtils.GbLogLevel = mainUtils.ERROR
		break
	}

	// If the file doesn't exist, create it, or append to the file
	mainUtils.Log("Starting Program", mainUtils.INFO)

	defer func() {
		if err = mainUtils.GpcFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	mainUtils.AcRouter = mux.NewRouter()
	mainUtils.AcRouter.HandleFunc("/home", handlers.GetHomeHandler).Methods("GET")
	mainUtils.AcRouter.HandleFunc("/SignOut", handlers.GetSignOutHandler).Methods("GET")
	mainUtils.AcRouter.HandleFunc("/closeWindow", handlers.GetCloseHandler).Methods("GET")
	mainUtils.AcRouter.HandleFunc("/writeExcel", handlers.GetExcelWriterHandler).Methods("GET")
	mainUtils.AcRouter.HandleFunc("/delete", handlers.PostDeleteHandler).Methods("POST")

	mainUtils.AcRouter.HandleFunc("/edit/{id:[0-9]+}", handlers.GetEditHandler).Methods("GET")
	mainUtils.AcRouter.HandleFunc("/editQty", handlers.PostEditQtyHandler).Methods("POST")

	mainUtils.AcRouter.HandleFunc("/editDrug", handlers.GetDrugEditHandler).Methods("GET")
	mainUtils.AcRouter.HandleFunc("/editDrug", handlers.PostDrugEditHandler).Methods("POST")

	mainUtils.AcRouter.HandleFunc("/editDrugGetNdc", handlers.GetDrugEditGetNdcHandler).Methods("GET")
	mainUtils.AcRouter.HandleFunc("/editDrugGetNdc", handlers.PostDrugEditGetNdcHandler).Methods("POST")

	mainUtils.AcRouter.HandleFunc("/delete/{id:[0-9]+}", handlers.GetDeleteHandler).Methods("GET")

	mainUtils.AcRouter.HandleFunc("/newDrug", handlers.GetNewDrugHandler).Methods("GET")
	mainUtils.AcRouter.HandleFunc("/newDrug", handlers.PostNewDrugHandler).Methods("POST")

	mainUtils.AcRouter.HandleFunc("/", handlers.GetDatabaseNdcHandler).Methods("GET")
	mainUtils.AcRouter.HandleFunc("/databaseDrug", handlers.GetDatabaseNdcHandler).Methods("GET")
	mainUtils.AcRouter.HandleFunc("/databaseDrug", handlers.PostDatabaseNdcHandler).Methods("POST")

	mainUtils.AcRouter.HandleFunc("/databaseDrug/{ndc:[0-9]{5}-[0-9]{4}-[0-9]{2}}",
		handlers.GetDatabaseNdcClickHandler).Methods("GET")

	mainUtils.AcRouter.HandleFunc("/databaseName", handlers.GetDatabaseNameHandler).Methods("GET")
	mainUtils.AcRouter.HandleFunc("/databaseName", handlers.PostDatabaseNameHandler).Methods("POST")

	mainUtils.AcRouter.HandleFunc("/audit", handlers.GetAuditHandler).Methods("GET")
	mainUtils.AcRouter.HandleFunc("/audit", handlers.PostAuditHandler).Methods("POST")

	mainUtils.AcRouter.HandleFunc("/purchase", handlers.GetPurchaseHandler).Methods("GET")
	mainUtils.AcRouter.HandleFunc("/purchase", handlers.PostPurchaseHandler).Methods("POST")

	mainUtils.AcRouter.HandleFunc("/prescription", handlers.GetPrescriptionHandler).Methods("GET")
	mainUtils.AcRouter.HandleFunc("/prescription", handlers.PostPrescriptionHandler).Methods("POST")

	// mainUtils.AcRouter.HandleFunc("/login", handlers.GetLoginHandler).Methods("GET")
	// mainUtils.AcRouter.HandleFunc("/login", handlers.PostLoginHandler).Methods("POST")

	// mainUtils.AcRouter.HandleFunc("/register", handlers.GetRegisterHandler).Methods("GET")
	// mainUtils.AcRouter.HandleFunc("/register", handlers.PostRegisterHandler).Methods("POST")

	http.Handle("/web/assets/", http.StripPrefix("/web/assets", http.FileServer(http.Dir("./web/assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", mainUtils.AcRouter)
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		mainUtils.Log("Failed to listen on port 80", mainUtils.ERROR)
		log.Fatal(err)
	}
}
