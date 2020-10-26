package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/jimlawless/whereami"
	"github.com/keighl/barkup"
)

func Backup() {

	postgres := &barkup.Postgres{
		Host: "127.0.0.1",
		Port: "5432",
		DB:   "drugrecord",

		// Not necessary if the program runs as an authorized pg user/role
		Username: "postgres",

		// Any extra pg_dump options
		Options: []string{"--no-owner"},
	}

	// Writes a file `./bu_DBNAME_TIMESTAMP.sql.tar.gz`
	result := postgres.Export()
	if result.Error != nil {
		Log(result.Error.Error(), ERROR, whereami.WhereAmI())
	} else {
		lcCurrentDate := time.Now().Format("01.02.2006")
		lcBackupName := fmt.Sprintf(`backup_%v.sql`, lcCurrentDate)

		lcPath := "."
		lacFiles, err := ioutil.ReadDir(lcPath)
		if err != nil {
			Log(err.Error(), ERROR, whereami.WhereAmI())
		}
		for _, lcFile := range lacFiles {
			if strings.Contains(lcFile.Name(), "bu_") {
				err = os.Rename(lcFile.Name(), "backups/"+lcBackupName)
				if err != nil {
					Log(err.Error(), ERROR, whereami.WhereAmI())
				}
			}
		}

		Log("Backup written to backups/"+lcBackupName, INFO, whereami.WhereAmI())
	}

	lcBackupPath := "backups/"
	lacFiles, err := ioutil.ReadDir(lcBackupPath)
	if err != nil {
		Log(err.Error(), ERROR, whereami.WhereAmI())
	}

	lcWeekAgo := time.Now().AddDate(0, 0, -7)
	for _, lcFile := range lacFiles {
		if lcFile.ModTime().AddDate(0, 0, -6).Before(lcWeekAgo) {
			if err = os.Remove("backups\\" + lcFile.Name()); err != nil {
				Log(err.Error(), ERROR, whereami.WhereAmI())
			}
		}
	}
	Log("Backup cleanup complete", INFO, whereami.WhereAmI())
}
