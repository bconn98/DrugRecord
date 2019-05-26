/**
File: userDB
Description: Works with the userdb
@author: Bryan Conn
@date: 10/7/18
*/
package mainUtils

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

/**
Function: issue
Description: Checks for an error and reports it
@param err The error
*/
func issue(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Database password needs to be changed when released
var connStr = "postgres://postgres:Zoo123@localhost/drugrecord?sslmode=disable"
var db, err = sql.Open("postgres", connStr)

/**
Function: GetUsers
Description: Grabs all of the users from the database
@return An array of User structs
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
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		err := rows.Scan(&lcUserName, &lnPassVal)
		issue(err)
		user := User{lcUserName, lnPassVal}
		users = append(users, user)
	}
	err = rows.Err()
	issue(err)

	return users
}

/**
Function: AddUser
Description: Adds a user to the database
@param acUsername The username of the new user
@param acPassVal The password value for the new user
*/
func AddUser(acUsername string, acPassVal int) {
	_, err := db.Query("INSERT INTO userDB (userName, passVal) VALUES ($1, $2);", acUsername, acPassVal)
	issue(err)
}
