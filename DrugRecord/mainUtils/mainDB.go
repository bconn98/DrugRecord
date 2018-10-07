package mainUtils

import (
	_ "github.com/lib/pq"
	"strconv"
)

func FindNDC(ndc string) ([]Order) {
	var NDC string
	var id int
	issue(err)

	rows, err := db.Query("SELECT id, ndc FROM orderdb WHERE ndc = $1;", ndc)
	issue(err)

	defer rows.Close()
	var orders []Order
	for rows.Next() {
		err := rows.Scan(&id, &NDC)
		issue(err)
		orders = append(orders, MakeOrder(id))
	}
	err = rows.Err()
	issue(err)

	return orders
}

func AddAudit(ndc string, pharmacist string, monthS string, dayS string, yearS string, qtyS string){
	month, _ := strconv.Atoi(monthS)
	day, _ := strconv.Atoi(dayS)
	year, _ := strconv.Atoi(yearS)
	qty, _ := strconv.Atoi(qtyS)
	_, err = db.Query("INSERT INTO orderdb (ndc, pharmacist, qty, date, logdate, script, audit, purchase, id) " +
		"VALUES ($1, $2, $3, make_date($4, $5, $6), current_date, $7, $8, $9, $10);", ndc, pharmacist, qty, year, month,
		day, false, true, false, 2)
}

func AddPrescription(ndc string, pharmacist string, monthS string, dayS string, yearS string, qtyS string){
	month, _ := strconv.Atoi(monthS)
	day, _ := strconv.Atoi(dayS)
	year, _ := strconv.Atoi(yearS)
	qty, _ := strconv.Atoi(qtyS)
	_, err = db.Query("INSERT INTO orderdb (ndc, pharmacist, qty, date, logdate, script, audit, purchase, id) " +
		"VALUES ($1, $2, $3, make_date($4, $5, $6), current_date, $7, $8, $9, $10);", ndc, pharmacist, qty, year, month,
		day, true, false, false, 3)
}

func AddPurchase(ndc string, pharmacist string, monthS string, dayS string, yearS string, qtyS string){
	month, _ := strconv.Atoi(monthS)
	day, _ := strconv.Atoi(dayS)
	year, _ := strconv.Atoi(yearS)
	qty, _ := strconv.Atoi(qtyS)
	_, err = db.Query("INSERT INTO orderdb (ndc, pharmacist, qty, date, logdate, script, audit, purchase, id) " +
		"VALUES ($1, $2, $3, make_date($4, $5, $6), current_date, $7, $8, $9, $10);", ndc, pharmacist, qty, year, month,
		day, false, false, true, 4)
}