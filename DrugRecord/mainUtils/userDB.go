package mainUtils

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
)

func issue(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var connStr = "postgres://postgres:Zoo123@localhost/DrugRecord?sslmode=disable"
var db, err = sql.Open("postgres", connStr)


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

func AddUser(username string, passVal int) {
	_, err := db.Query("INSERT INTO userDB (userName, passVal) VALUES ($1, $2);", username, passVal)
	issue(err)
}
