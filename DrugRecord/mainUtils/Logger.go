package mainUtils

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var GpcFile *os.File
var Initial bool

func LogSql(acEvent string) {
	updateLogStructure()
	_, err := GpcFile.WriteString(acEvent + "\n")
	if err != nil {
		log.Fatal(err)
	}
}

func LogError(acError string) {
	if _, err := GpcFile.WriteString("ERROR: " + acError + "\n"); err != nil {
		log.Fatal(err)
	}
}

func updateLogStructure() {
	var fileName string
	currentTimeString := time.Now().Format("01-2006")
	currentTimeString = strings.Replace(currentTimeString, "-", "_", -1)
	fileName = "log/" + currentTimeString + "_log.log"
	if !fileExists(fileName) {
		if GpcFile != nil {
			if err = GpcFile.Close(); err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println(fileName)
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
