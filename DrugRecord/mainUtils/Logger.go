package mainUtils

import (
	"fmt"
	"io"
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

func GetLastNdc() string {
	var lcNdc string

	lastLine := getLastLineWithSeek()

	runes := []rune(lastLine)

	// ... Convert back into a string from rune slice.
	lcNdc = string(runes[0:14])
	fmt.Println(lcNdc)

	return lcNdc
}

func getLastLineWithSeek() string {
	line := ""
	var cursor int64 = 0
	stat, _ := GpcFile.Stat()
	filesize := stat.Size()
	for {
		cursor -= 1
		GpcFile.Seek(cursor, io.SeekEnd)

		char := make([]byte, 1)
		GpcFile.Read(char)

		if cursor != -1 && (char[0] == 10 || char[0] == 13) { // stop if we find a line
			break
		}

		line = fmt.Sprintf("%s%s", string(char), line) // there is more efficient way

		if cursor == -filesize { // stop if we are at the begining
			break
		}
	}

	return line
}
