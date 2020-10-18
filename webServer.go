/**
File: webServer
Description: Runs a database
@author Bryan Conn
@date 10/7/2018
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gorilla/mux"
	"gopkg.in/go-ini/ini.v1"

	"github.com/bconn98/DrugRecord/utils"
	"github.com/bconn98/DrugRecord/web/handlers"
)

func test() {
	fmt.Println("Fuck off")
}

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
	utils.Log("Starting Program", utils.INFO)

	lcGoScheduler := gocron.NewScheduler(time.UTC)
	_, err = lcGoScheduler.Every(1).Day().At("04:00").Do(utils.Backup)
	if err != nil {
		utils.Log(err.Error(), utils.ERROR)
	}
	lcGoScheduler.StartAsync()

	defer func() {
		if err = utils.GpcFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	utils.AcRouter = mux.NewRouter()
	utils.AcRouter.HandleFunc("/home", handlers.GetHomeHandler).Methods("GET")
	utils.AcRouter.HandleFunc("/SignOut", handlers.GetSignOutHandler).Methods("GET")
	utils.AcRouter.HandleFunc("/closeWindow", handlers.GetCloseHandler).Methods("GET")
	utils.AcRouter.HandleFunc("/writeExcel", handlers.GetExcelWriterHandler).Methods("GET")
	utils.AcRouter.HandleFunc("/delete", handlers.PostDeleteHandler).Methods("POST")

	utils.AcRouter.HandleFunc("/edit/{id:[0-9]+}", handlers.GetEditHandler).Methods("GET")
	utils.AcRouter.HandleFunc("/editQty", handlers.PostEditQtyHandler).Methods("POST")

	utils.AcRouter.HandleFunc("/editDrug", handlers.GetDrugEditHandler).Methods("GET")
	utils.AcRouter.HandleFunc("/editDrug", handlers.PostDrugEditHandler).Methods("POST")

	utils.AcRouter.HandleFunc("/editDrugGetNdc", handlers.GetDrugEditGetNdcHandler).Methods("GET")
	utils.AcRouter.HandleFunc("/editDrugGetNdc", handlers.PostDrugEditGetNdcHandler).Methods("POST")

	utils.AcRouter.HandleFunc("/delete/{id:[0-9]+}", handlers.GetDeleteHandler).Methods("GET")

	utils.AcRouter.HandleFunc("/newDrug", handlers.GetNewDrugHandler).Methods("GET")
	utils.AcRouter.HandleFunc("/newDrug", handlers.PostNewDrugHandler).Methods("POST")

	utils.AcRouter.HandleFunc("/", handlers.GetDatabaseNdcHandler).Methods("GET")
	utils.AcRouter.HandleFunc("/databaseDrug", handlers.GetDatabaseNdcHandler).Methods("GET")
	utils.AcRouter.HandleFunc("/databaseDrug", handlers.PostDatabaseNdcHandler).Methods("POST")

	utils.AcRouter.HandleFunc("/databaseDrug/{ndc:[0-9]{5}-[0-9]{4}-[0-9]{2}}",
		handlers.GetDatabaseNdcClickHandler).Methods("GET")

	utils.AcRouter.HandleFunc("/databaseName", handlers.GetDatabaseNameHandler).Methods("GET")
	utils.AcRouter.HandleFunc("/databaseName", handlers.PostDatabaseNameHandler).Methods("POST")

	utils.AcRouter.HandleFunc("/audit", handlers.GetAuditHandler).Methods("GET")
	utils.AcRouter.HandleFunc("/audit", handlers.PostAuditHandler).Methods("POST")

	utils.AcRouter.HandleFunc("/purchase", handlers.GetPurchaseHandler).Methods("GET")
	utils.AcRouter.HandleFunc("/purchase", handlers.PostPurchaseHandler).Methods("POST")

	utils.AcRouter.HandleFunc("/prescription", handlers.GetPrescriptionHandler).Methods("GET")
	utils.AcRouter.HandleFunc("/prescription", handlers.PostPrescriptionHandler).Methods("POST")

	// mainUtils.AcRouter.HandleFunc("/login", handlers.GetLoginHandler).Methods("GET")
	// mainUtils.AcRouter.HandleFunc("/login", handlers.PostLoginHandler).Methods("POST")

	// mainUtils.AcRouter.HandleFunc("/register", handlers.GetRegisterHandler).Methods("GET")
	// mainUtils.AcRouter.HandleFunc("/register", handlers.PostRegisterHandler).Methods("POST")

	http.Handle("/web/assets/", http.StripPrefix("/web/assets", http.FileServer(http.Dir("./web/assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", utils.AcRouter)
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		utils.Log("Failed to listen on port 80", utils.ERROR)
		log.Fatal(err)
	}
}
