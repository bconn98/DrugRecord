package mainUtils

import (
	"log"
	"os"
	"strings"
	"time"
)

type LogLevel int

const (
	DEBUG   LogLevel = 0
	SQL              = 1
	INFO             = 2
	WARNING          = 3
	ERROR            = 4
)

var GpcFile *os.File
var Initial bool
var GbLogLevel LogLevel

func Log(acLog string, anLogLevel LogLevel) {
	updateLogStructure()
	switch anLogLevel {
	case DEBUG:
		logAll(acLog, "DEBUG", anLogLevel)
		break
	case SQL:
		logAll(acLog, "SQL", anLogLevel)
		break
	case INFO:
		logAll(acLog, "INFO", anLogLevel)
		break
	case WARNING:
		logAll(acLog, "WARNING", anLogLevel)
		break
	case ERROR:
		logAll(acLog, "ERROR", anLogLevel)
		break
	}
}

func logAll(acLog string, acLogLevel string, anLogLevel LogLevel) {
	currentDateString := time.Now().Format("2006-01-02 15:04:05")
	if GbLogLevel <= anLogLevel {
		if _, err := GpcFile.WriteString(currentDateString + " " + acLogLevel + " " + acLog + "\n"); err != nil {
			log.Fatal(err)
		}
	}
}

func updateLogStructure() {
	var fileName string
	currentDateString := time.Now().Format("01-2006")
	currentDateString = strings.Replace(currentDateString, "-", "_", -1)
	fileName = "log/" + currentDateString + "_log.log"
	if !fileExists(fileName) {
		if GpcFile != nil {
			if err = GpcFile.Close(); err != nil {
				log.Fatal(err)
			}
		}

		GpcFile, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	if Initial {
		Initial = false
		GpcFile, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
