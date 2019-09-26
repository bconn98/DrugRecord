package mainUtils

import (
	"log"
	"os"
)

var GpcFile *os.File

func LogSql( acEvent string ) {
	_, err := GpcFile.WriteString(acEvent + "\n")
	if err != nil {
		log.Fatal(err)
	}
}

func LogError( acError string ) {
	_, err := GpcFile.WriteString("ERROR: " + acError + "\n")
	if err != nil {
		log.Fatal(err)
	}
}
