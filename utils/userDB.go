/**
File: userDB
Description: Works with the userdb
@author: Bryan Conn
@date: 10/7/18
*/
package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

/**
Function: issue
Description: Checks for an error and reports it
@param err The error
*/
func issue(err error) {
	if err != nil {
		Log(err.Error(), ERROR)
	}
}

// Database password needs to be changed when released
var connStr = "postgres://postgres:Zoo123@localhost/drugrecord?sslmode=disable"
var db, err = sql.Open("postgres", connStr)

/**
Function: GetUsers
Description: Grabs all of the users from the database
@return An array of User structures
*/
func GetUsers() []User {
	var (
		lcUserName string
		lnPassVal  int
	)
	issue(err)

	rows, err := db.Query("SELECT * FROM userDB;")
	issue(err)

	var users []User

	for rows.Next() {
		if rows.Err() != nil {
			issue(rows.Err())
			break
		}
		issue(rows.Scan(&lcUserName, &lnPassVal))
		users = append(users, User{lcUserName, lnPassVal})
	}

	defer func() {
		issue(rows.Close())
	}()

	return users
}

/**
Function: AddUser
Description: Adds a user to the database
@param acUsername The username of the new user
@param anPassVal The password value for the new user
*/
func AddUser(acUsername string, anPassVal int) {
	insertString := fmt.Sprintf("%s%s%s%d%s", "INSERT INTO userdb (userName, passVal) VALUES ('", acUsername,
		"', ", anPassVal, ");")
	_, err := db.Exec(insertString)
	issue(err)
	Log(insertString, SQL)
}
