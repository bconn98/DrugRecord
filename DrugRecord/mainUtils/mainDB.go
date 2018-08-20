package mainUtils

import (
	_ "github.com/lib/pq"
	"strconv"
	"fmt"
)

func FindUDC(udc string) ([]string) {
	var UDC string
	issue(err)

	rows, err := db.Query("SELECT udc FROM mainDB WHERE udc = $1;", udc)
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

func AddAudit(udc string, pharmacist string, monthS string, dayS string,
	yearS string, amonthS string, adayS string, ayearS string, qtyS string){
	ldate := insertDate(monthS, dayS, yearS)
	adate := insertDate(amonthS, adayS, ayearS)
	qty, _ := strconv.Atoi(qtyS)
	_, err = db.Query("INSERT INTO maindb (udc, pharmacist, aqty, adate, ldate) VALUES ($1, $2, $3, $4, $5);",
		udc, pharmacist, qty, adate, ldate)
}

func insertDate(monthS string, dayS string, yearS string) (Date){
	var d1 Date
	month, _ := strconv.Atoi(monthS)
	day, _ := strconv.Atoi(dayS)
	year, _ := strconv.Atoi(yearS)
	testString := fmt.Sprintf("DO $$ BEGIN DECLARE d1 \"date\"; d1 := (%d, %d, %d); RETURNING d1; END $$",
		month, day, year)
	err := db.QueryRow(testString).Scan(&d1)
	issue(err)
	return d1
}

func AddPrescription(){

}

func AddPurchase(){

}