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
func FindNDC(acNdc string) (string, string, string, string, string, time.Time, float64, []Order) {
	var lcPharmacist string
	var lcScript string
	var lcDate time.Time
	var lnQty float64
	var lcType string
	var lcName string
	var lcForm string
	var lcItemNum string
	var lcSize string
	var lnDrugQty float64
	var lasOrders []Order
	rows, err := db.Query(
		"SELECT pharmacist, date, qty, script, type FROM orderdb WHERE ndc = $1 order by date desc;", acNdc)
	issue(err)

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		err := rows.Scan(&lcPharmacist, &lcDate, &lnQty, &lcScript, &lcType)
		issue(err)
		lasOrders = append(lasOrders, MakeOrder(lcPharmacist, lcScript, lcType, lnQty, lcDate))
	}
	err = rows.Err()
	issue(err)
	row, err := db.Query("Select name, form, item_num, size, date, qty from drugdb where ndc = $1", acNdc)
	issue(err)
	row.Next()
	err = row.Scan(&lcName, &lcForm, &lcItemNum, &lcSize, &lcDate, &lnDrugQty)
	issue(err)
	return lcName, acNdc, lcForm, lcItemNum, lcSize, lcDate, lnDrugQty, lasOrders
}

/**
Function: addType
Description: Creates a new row in the orderdb with passed in attributes
@param acNdc The NDC of the drug of interest
@param acPharmacist The pharmacist who is inputting the order
@param acMonth The month the order was made
@param anDay The day the order was made
@param anYear The year the order was made
@param anQty The quantity of the order
@param acScript The order#/Script# or blank
@param acOrderType The type of the order
*/
func addType(acNdc string, acPharmacist string, acMonth string, anDay string, anYear string,
	anQty string, acScript string, acOrderType string) bool {
	lnMonth, _ := strconv.Atoi(acMonth)
	lnDay, _ := strconv.Atoi(anDay)
	lnYear, _ := strconv.Atoi(anYear)
	qty, _ := strconv.ParseFloat(anQty, 64)

	row, err := db.Query("Select count(script) from orderdb where script = $1 and "+
		"date = make_date($2, $3, $4) and qty = $5 and ndc = $6;", acScript, lnYear, lnMonth, lnDay, qty, acNdc)

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
		"VALUES ($1, $2, $3, make_date($4, $5, $6), current_date, $7, $8);", acNdc, acPharmacist, qty, lnYear, lnMonth,
		lnDay, acScript, acOrderType)

	return true
}

/**
Function: alterQty
Description: Alters a drugs quantity using its NDC to find it
@param acNdc The ndc value of the drug in question
@param acQty The quantity of the alteration
*/
func alterQty(acNdc string, acQty string) {
	var lnRowQty float64

	if strings.Contains(acQty, "--") {
		acQty = strings.Replace(acQty, "--", "-", 1)
	}

	lnQty, _ := strconv.ParseFloat(acQty, 64)
	rows, err := db.Query("SELECT qty from drugdb where ndc = $1", acNdc)
	issue(err)
	rows.Next()
	err = rows.Scan(&lnRowQty)
	if err != nil {
		_, err = db.Query("insert into drugdb (ndc, qty) values ($1, $2)", acNdc, lnQty)
		return
	}
	lnNewQty := lnRowQty - lnQty
	_, err = db.Query("UPDATE drugdb set qty = $1 where ndc = $2", lnNewQty, acNdc)
}

/**
Function: setDrugQty
Description: Sets the drug quantity to the new value
@param acNdc The ndc value of the drug in question
@param acQty The new qty of the drug
@return The difference in change
*/
func setDrugQty(acNdc string, acQty string) int {
	var lnRowQty int
	rows, err := db.Query("SELECT qty from drugdb where ndc = $1", acNdc)
	issue(err)
	rows.Next()
	err = rows.Scan(&lnRowQty)
	lnQty, _ := strconv.Atoi(acQty)
	_, err = db.Query("UPDATE drugdb set qty = $1 where ndc = $2", lnQty, acNdc)
	return lnRowQty - lnQty
}

/**
Function: AddPrescription
Description: Adds a prescription type order to the orderdb
@param acNdc The NDC of the drug of interest
@param acPharmacist The pharmacist who is inputting the prescription
@param acMonth The month the prescription was made
@param acDay The day the prescription was made
@param acYear The year the prescription was made
@param acQty The quantity of the prescription
@param acScript The prescription number
@param acActual The actual count if entered
*/
func AddPrescription(acNdc string, acPharmacist string, acMonth string, acDay string,
	acYear string, acQty string, acScript string, acActual string) bool {

	var lbCheck bool
	lbCheck = addType(acNdc, acPharmacist, acMonth, acDay, acYear, acQty, acScript, "Prescription")
	if !lbCheck {
		return false
	}

	alterQty(acNdc, acQty)

	if acActual != "" {
		lnQtyDiff := setDrugQty(acNdc, acActual)
		lbCheck = addType(acNdc, acPharmacist, acMonth, acDay, acYear,
			strconv.Itoa(lnQtyDiff), acScript, "Actual Count")
	}
	return lbCheck
}

/**
Function: AddAudit
Description: Adds a audit type order to the orderdb
@param acNdc The NDC of the drug of interest
@param acPharmacist The pharmacist who is inputting the audit
@param acMonth The month the audit was made
@param acDay The day the audit was made
@param acYear The year the audit was made
@param acQty The quantity of the audit
@param acActual The actual count if entered
*/
func AddAudit(acNdc string, acPharmacist string, acMonth string, acDay string,
	acYear string, acQty string, acActual string) bool {
	var lbCheck bool
	lbCheck = addType(acNdc, acPharmacist, acMonth, acDay, acYear, acQty, "", "Audit")
	if !lbCheck {
		return false
	}

	if acActual != "" {
		lnQtyDiff := setDrugQty(acNdc, acActual)
		lbCheck = addType(acNdc, acPharmacist, acMonth, acDay, acYear,
			strconv.Itoa(lnQtyDiff), "", "Actual Count")
	}
	return lbCheck
}

/**
Function: AddPurchase
Description: Adds a purchase type order to the orderdb
@param acNdc The NDC of the drug of interest
@param acPharmacist The pharmacist who is inputting the purchase
@param acMonth The month the purchase was made
@param acDay The day the purchase was made
@param acYear The year the purchase was made
@param acQty The quantity of the purchase
@param acActual The actual count if entered
*/
func AddPurchase(acNdc string, acPharmacist string, acMonth string, acDay string,
	acYear string, acQty string, order string, acActual string) bool {

	var lbCheck bool
	lbCheck = addType(acNdc, acPharmacist, acMonth, acDay, acYear, acQty, order, "Purchase")

	if !lbCheck {
		return false
	}
	alterQty(acNdc, "-"+acQty)

	if acActual != "" {
		lnQtyDiff := setDrugQty(acNdc, acActual)
		lbCheck = addType(acNdc, acPharmacist, acMonth, acDay, acYear,
			strconv.Itoa(lnQtyDiff), order, "Actual Count")
	}
	return lbCheck
}

/**
Function: NewCheck
Description: See if the drug is in the database yet
@param acNdc The NDC of the drug to check
@return If the drug is in the database
*/
func NewCheck(acNdc string) bool {
	var lnCount int
	row, err := db.Query("Select count(ndc) from drugdb where ndc = $1", acNdc)
	issue(err)
	row.Next()
	err = row.Scan(&lnCount)
	issue(err)

	if lnCount < 1 {
		return false
	} else {
		return true
	}
}

/**
Function: AddDrug
Description: Adds the drug to the database without the defaults
@param acNdc The ndc
@param acMonth The month as a string
@param acDay The day as a string
@param acYear The year as a string
*/
func AddDrug(acNdc string, acMonth string, acDay string, acYear string) {
	lnMonth, _ := strconv.Atoi(acMonth)
	lnDay, _ := strconv.Atoi(acDay)
	lnYear, _ := strconv.Atoi(acYear)
	_, err = db.Query("Insert into drugdb (ndc, date) values ($1, make_date($2, $3, $4))",
		acNdc, lnYear, lnMonth, lnDay)
	issue(err)
}

/**
Function: UpdateDrug
Description: Adds the correct default values to a drug
@param acSize The size of the packet
@param acForm The form of the drug
@param acItemNum The item number of the drug
@param acName The name of the drug
@param acNdc The ndc of the drug
*/
func UpdateDrug(acSize string, acForm string, acItemNum string, acName string, acNdc string) {
	_, _ = db.Query("Update drugdb set size = $1, form = $2, item_num = $3, name = $4 where ndc = $5",
		acSize, acForm, acItemNum, acName, acNdc)
}
