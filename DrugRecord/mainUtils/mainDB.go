/**
File: mainDB
Description: Does the all of the work with the order database
@author: Bryan Conn
@date: 10/7/18
 */
package mainUtils

import (
	_ "github.com/lib/pq"
	"strconv"
)

/**
Function: FindNDC
Description: Given a NDC, finds and returns all corresponding rows
@param ndc The NDC to match
@return An array of orders with the given NDC
 */
func FindNDC(ndc string) ([]Order) {
	var NDC string
	var id int
	issue(err)

	rows, err := db.Query("SELECT id, ndc FROM orderdb WHERE ndc = $1;", ndc)
	issue(err)

	defer rows.Close()
	//Find a way to easily have all the data
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

/**
Function: addType
Description: Creates a new row in the orderdb with passed in attributes
@param ndc The NDC of the drug of interest
@param pharmacist The pharmacist who is inputting the order
@param monthS The month the order was made
@param dayS The day the order was made
@param yearS The year the order was made
@param qtyS The quantity of the order
@param typ An int value to determine the type of the order
 */
func addType(ndc string, pharmacist string, monthS string, dayS string, yearS string, qtyS string, typ int) {
	var script, audit, purchase bool
	switch typ  {
		case 0:
			script = true
			audit = false
			purchase = false
		case 1:
			script = false
			audit = true
			purchase = false
		case 2:
			script = false
			audit = false
			purchase = true
	}
	month, _ := strconv.Atoi(monthS)
	day, _ := strconv.Atoi(dayS)
	year, _ := strconv.Atoi(yearS)
	qty, _ := strconv.Atoi(qtyS)
	_, err = db.Query("INSERT INTO orderdb (ndc, pharmacist, qty, date, logdate, script, audit, purchase, id) " +
		"VALUES ($1, $2, $3, make_date($4, $5, $6), current_date, $7, $8, $9, $10);", ndc, pharmacist, qty, year, month,
		day, script, audit, purchase, 2)
}

/**
Function: AddPrescription
Description: Adds a prescription type order to the orderdb
@param ndc The NDC of the drug of interest
@param pharmacist The pharmacist who is inputting the prescription
@param monthS The month the prescription was made
@param dayS The day the prescription was made
@param yearS The year the prescription was made
@param qtyS The quantity of the prescription
 */
func AddPrescription(ndc string, pharmacist string, monthS string, dayS string, yearS string, qtyS string){
	addType(ndc, pharmacist, monthS, dayS, yearS, qtyS, 0)
}

/**
Function: AddAudit
Description: Adds a audit type order to the orderdb
@param ndc The NDC of the drug of interest
@param pharmacist The pharmacist who is inputting the audit
@param monthS The month the audit was made
@param dayS The day the audit was made
@param yearS The year the audit was made
@param qtyS The quantity of the audit
 */
func AddAudit(ndc string, pharmacist string, monthS string, dayS string, yearS string, qtyS string){
	addType(ndc, pharmacist, monthS, dayS, yearS, qtyS, 1)
}

/**
Function: AddPurchase
Description: Adds a purchase type order to the orderdb
@param ndc The NDC of the drug of interest
@param pharmacist The pharmacist who is inputting the purchase
@param monthS The month the purchase was made
@param dayS The day the purchase was made
@param yearS The year the purchase was made
@param qtyS The quantity of the purchase
 */
func AddPurchase(ndc string, pharmacist string, monthS string, dayS string, yearS string, qtyS string) {
	addType(ndc, pharmacist, monthS, dayS, yearS, qtyS, 2)
}