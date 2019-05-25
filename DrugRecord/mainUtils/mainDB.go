/**
File: mainDB
Description: Does the all of the work with the order database
@author: Bryan Conn
@date: 10/7/18
*/
package mainUtils

import (
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"strings"
	"time"
)

/**
Function: FindNDC
Description: Given a NDC, finds and returns all corresponding rows
@param ndc The NDC to match
@return An array of orders with the given NDC, vital drug information
*/
func FindNDC(ndc string) (string, string, string, string, string, time.Time, float64, []Order) {
	var NDC string
	var pharm string
	var script string
	var date time.Time
	var qty float64
	var typ string
	var name string
	var form string
	var itemNum string
	var size string
	var drugQty float64
	var orders []Order
	rows, err := db.Query("SELECT ndc, pharmacist, date, qty, script, type FROM orderdb WHERE ndc = $1 order by date desc;", ndc)
	issue(err)

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		err := rows.Scan(&NDC, &pharm, &date, &qty, &script, &typ)
		issue(err)
		orders = append(orders, MakeOrder(pharm, script, typ, qty, date))
	}
	err = rows.Err()
	issue(err)
	row, err := db.Query("Select name, ndc, form, item_num, size, date, qty from drugdb where ndc = $1", ndc)
	issue(err)
	row.Next()
	err = row.Scan(&name, &NDC, &form, &itemNum, &size, &date, &drugQty)
	issue(err)
	return name, NDC, form, itemNum, size, date, drugQty, orders
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
@param script The order#/Script# or blank
@param orderType The type of the order
*/
func addType(ndc string, pharmacist string, monthS string, dayS string, yearS string,
	qtyS string, script string, orderType string) bool {
	month, _ := strconv.Atoi(monthS)
	day, _ := strconv.Atoi(dayS)
	year, _ := strconv.Atoi(yearS)
	qty, _ := strconv.ParseFloat(qtyS, 64)

	row, err := db.Query("Select count(script) from orderdb where script = $1 and "+
		"date = make_date($2, $3, $4) and qty = $5 and ndc = $6;", script, year, month, day, qty, ndc)

	if err != nil {
		issue(err)
	}

	count := 0
	row.Next()
	err = row.Scan(&count)

	if err != nil {
		issue(err)
	}

	if count != 0 {
		return false
	}

	_, err = db.Query("INSERT INTO orderdb (ndc, pharmacist, qty, date, logdate, script, type) "+
		"VALUES ($1, $2, $3, make_date($4, $5, $6), current_date, $7, $8);", ndc, pharmacist, qty, year, month,
		day, script, orderType)

	return true
}

/**
Function: alterQty
Description: Alters a drugs quantity using its NDC to find it
@param ndc The ndc value of the drug in question
@param qtyS The quantity of the alteration
*/
func alterQty(ndc string, qtyS string) {
	var rowQ float64

	if strings.Contains(qtyS, "--") {
		qtyS = strings.Replace(qtyS, "--", "-", 1)
	}

	qty, _ := strconv.ParseFloat(qtyS, 64)
	rows, err := db.Query("SELECT qty from drugdb where ndc = $1", ndc)
	issue(err)
	rows.Next()
	err = rows.Scan(&rowQ)
	if err != nil {
		_, err = db.Query("insert into drugdb (ndc, qty) values ($1, $2)", ndc, qty)
		return
	}
	newQ := rowQ - qty
	_, err = db.Query("UPDATE drugdb set qty = $1 where ndc = $2", newQ, ndc)
}

/**
Function: setDrugQty
Description: Sets the drug quantity to the new value
@param ndc The ndc value of the drug in question
@param qtyS The new qty of the drug
@return The difference in change
*/
func setDrugQty(ndc string, qtyS string) int {
	var rowQ int
	rows, err := db.Query("SELECT qty from drugdb where ndc = $1", ndc)
	issue(err)
	rows.Next()
	err = rows.Scan(&rowQ)
	qty, _ := strconv.Atoi(qtyS)
	_, err = db.Query("UPDATE drugdb set qty = $1 where ndc = $2", qty, ndc)
	return rowQ - qty
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
@param script The prescription number
@param actualS The actual count if entered
*/
func AddPrescription(ndc string, pharmacist string, monthS string, dayS string,
	yearS string, qtyS string, script string, actualS string) bool {

	var check bool
	check = addType(ndc, pharmacist, monthS, dayS, yearS, qtyS, script, "Prescription")
	if !check {
		return false
	}

	alterQty(ndc, qtyS)

	if actualS != "" {
		diff := setDrugQty(ndc, actualS)
		check = addType(ndc, pharmacist, monthS, dayS, yearS, strconv.Itoa(diff), script, "Actual Count")
	}
	return check
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
@param actualS The actual count if entered
*/
func AddAudit(ndc string, pharmacist string, monthS string,
	dayS string, yearS string, qtyS string, actualS string) bool {
	var check bool
	check = addType(ndc, pharmacist, monthS, dayS, yearS, qtyS, "", "Audit")
	if !check {
		return false
	}

	if actualS != "" {
		diff := setDrugQty(ndc, actualS)
		check = addType(ndc, pharmacist, monthS, dayS, yearS, strconv.Itoa(diff), "", "Actual Count")
	}
	return check
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
@param actualS The actual count if entered
*/
func AddPurchase(ndc string, pharmacist string, monthS string, dayS string,
	yearS string, qtyS string, order string, actualS string) bool {

	var check bool
	check = addType(ndc, pharmacist, monthS, dayS, yearS, qtyS, order, "Purchase")

	if !check {
		return false
	}
	alterQty(ndc, "-"+qtyS)

	if actualS != "" {
		diff := setDrugQty(ndc, actualS)
		check = addType(ndc, pharmacist, monthS, dayS, yearS, strconv.Itoa(diff), order, "Actual Count")
	}
	return check
}

/**
Function: NewCheck
Description: See if the drug is in the database yet
@param ndc The NDC of the drug to check
@return If the drug is in the database
*/
func NewCheck(ndc string) bool {
	var count int
	row, err := db.Query("Select count(ndc) from drugdb where ndc = $1", ndc)
	issue(err)
	row.Next()
	err = row.Scan(&count)
	issue(err)

	if count < 1 {
		return false
	} else {
		return true
	}
}

/**
Function: AddDrug
Description: Adds the drug to the database without the defaults
@param ndc The ndc
@param monthS The month as a string
@param dayS The day as a string
@param yearS The year as a string
*/
func AddDrug(ndc string, monthS string, dayS string, yearS string) {
	month, _ := strconv.Atoi(monthS)
	day, _ := strconv.Atoi(dayS)
	year, _ := strconv.Atoi(yearS)
	_, err = db.Query("Insert into drugdb (ndc, date) values ($1, make_date($2, $3, $4))",
		ndc, year, month, day)
	issue(err)
}

/**
Function: UpdateDrug
Description: Adds the correct default values to a drug
@param size The size of the packet
@param form The form of the drug
@param itemNum The item number of the drug
@param name The name of the drug
@param ndc The ndc of the drug
*/
func UpdateDrug(size string, form string, itemNum string, name string, ndc string) {
	_, _ = db.Query("Update drugdb set size = $1, form = $2, item_num = $3, name = $4 where ndc = $5",
		size, form, itemNum, name, ndc)
}
