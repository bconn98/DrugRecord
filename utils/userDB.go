/**
File: userDB
Description: Works with the userdb
@author: Bryan Conn
@date: 10/7/18
*/
package utils

import (
	"fmt"

	"github.com/jimlawless/whereami"
	_ "github.com/lib/pq"
)

/**
Function: issue
Description: Checks for an error and reports it
@param err The error
*/
func issue(err error, acPath string) {
	if err != nil {
		Log(err.Error(), ERROR, acPath)
	}
}

/**
Function: GetUsers
Description: Grabs all of the users from the database
@return An array of User structures
*/
func GetUsers() []User {
	var (
		lcUserName string
		lcPassVal  string
	)

	rows, err := McDb.Query("SELECT * FROM userDB;")
	issue(err, whereami.WhereAmI())

	var users []User
	// var testU = User{"Bryan", "AAA123"}
	// users = append(users, testU)

	for rows.Next() {
		if rows.Err() != nil {
			issue(rows.Err(), whereami.WhereAmI())
			break
		}
		issue(rows.Scan(&lcUserName, &lcPassVal), whereami.WhereAmI())
		users = append(users, User{lcUserName, lcPassVal})
	}

	defer func() {
		issue(rows.Close(), whereami.WhereAmI())
	}()

	return users
}

/**
Function: AddUser
Description: Adds a user to the database
@param acUsername The username of the new user
@param anPassVal The password value for the new user
*/
func AddUser(acUsername string, acPassVal string) {
	insertString := fmt.Sprintf("%s%s%s%s%s", "INSERT INTO userdb (userName, passVal) VALUES ('", acUsername,
		"', ", acPassVal, ");")
	_, err := McDb.Exec(insertString)
	issue(err, whereami.WhereAmI())
	Log(insertString, SQL, whereami.WhereAmI())
}
