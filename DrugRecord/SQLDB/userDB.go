package SQLDB

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"../mainUtils"
)

func issue(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetUsers() ([]mainUtils.User){
	var (
		userName string
		passVal int
	)
	connStr := "postgres://postgres:Zoo123@localhost/DrugRecord?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	issue(err)

	rows, err := db.Query("SELECT * FROM userDB")
	issue(err)

	var users []mainUtils.User
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&userName, &passVal)
		issue(err)
		user := mainUtils.User{userName, passVal}
		users = append(users, user)
	}
	err = rows.Err()
	issue(err)
	defer db.Close()
	return users
}