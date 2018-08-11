package mainUtils

import (
	_ "github.com/lib/pq"
)

func FindUDC(udc string) ([]string){
	var (
		UDC string
	)
	issue(err)

	rows, err := db.Query("SELECT UDC FROM mainDB WHERE UDC = $1;", udc)
	issue(err)

	defer rows.Close()
	var orders []string
	for rows.Next() {
		err := rows.Scan(&UDC)
		issue(err)
		orders = append(orders, UDC)
	}
	err = rows.Err()
	issue(err)

	return orders
}
