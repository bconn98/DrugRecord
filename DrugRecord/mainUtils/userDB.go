/**
File: userDB
Description: Works with the userdb
@author: Bryan Conn
@date: 10/7/18
 */
package mainUtils

import (
	_ "github.com/lib/pq"
	"log"
	"database/sql"
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
var connStr = "postgres://postgres:Zoo123@localhost/DrugRecord?sslmode=disable"
var db, err = sql.Open("postgres", connStr)

/**
Function: GetUsers
Description: Grabs all of the users from the database
@return An array of User structs
 */
func GetUsers() ([]User){
	var (
		userName string
		passVal int
	)
	issue(err)

	rows, err := db.Query("SELECT * FROM userDB;")
	issue(err)

	var users []User
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&userName, &passVal)
		issue(err)
		user := User{userName, passVal}
		users = append(users, user)
	}
	err = rows.Err()
	issue(err)

	return users
}

/**
Function: AddUser
Description: Adds a user to the database
@param username The username of the new user
@param passVal The password value for the new user
 */
func AddUser(username string, passVal int) {
	_, err := db.Query("INSERT INTO userDB (userName, passVal) VALUES ($1, $2);", username, passVal)
	issue(err)
}
