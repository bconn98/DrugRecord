/**
File: mainDB
Description: Does the all of the work with the order database
@author: Bryan Conn
@date: 10/7/18
*/
package mainUtils

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

/**
Function: FindNDC
Description: Given a NDC, finds AND returns all corresponding rows
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
	var lnId int64
	rows, err := db.Query(
		"SELECT pharmacist, date, qty, script, type, id FROM orderdb WHERE ndc = $1 ORDER BY date desc;", acNdc)
	issue(err)
	LogSql(fmt.Sprintf("%s%s%s", "SELECT pharmacist, date, qty, script, type, id FROM orderdb WHERE ndc = '", acNdc,
		"' ORDER BY date desc;"))

	for rows.Next() {
		err := rows.Scan(&lcPharmacist, &lcDate, &lnQty, &lcScript, &lcType, &lnId)
		issue(err)
		lasOrders = append(lasOrders, MakeOrder(acNdc, lcPharmacist, lcScript, lcType, lnQty, lcDate, lnId))
	}
	err = rows.Err()
	issue(err)

	row, err := db.Query("SELECT name, form, item_num, size, date, qty FROM drugdb WHERE ndc = $1", acNdc)
	issue(err)
	LogSql(fmt.Sprintf( "%s%s%s", "SELECT name, form, item_num, size, date, qty FROM drugdb WHERE ndc = '", acNdc, "';"))

	row.Next()
	err = row.Scan(&lcName, &lcForm, &lcItemNum, &lcSize, &lcDate, &lnDrugQty)

	if err != nil {
		return "", "", "", "", "", time.Time{}, 0, nil
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	return lcName, acNdc, lcForm, lcItemNum, lcSize, lcDate, lnDrugQty, lasOrders
}

/**
Function: GetOrder
Description: Gets the fields of an order that weren't specified by user
@param acNdc The ndc of the drug to get
@param acPharmacist The pharmacist who logged the order
@param acMonth The month the order was logged
@param acDay The day the order was logged
@param acYear The year the order was logged
@param acScript The script number of the order
@param acType The type of the order
*/
func GetOrder(acNdc string, acPharmacist string, acMonth string, acDay string, acYear string,
	acScript string, acType string) []Order {
	var lasOrder []Order
	var lnQty float64
	var rows *sql.Rows
	var err error
	var lcType string
	var lcDate time.Time
	var lnId int64
	if acType == "Audit" {
		rows, err = db.Query("SELECT qty, type, date, script, id FROM orderdb WHERE ndc = $1 AND pharmacist = $2 AND " +
			"date = make_date($3, $4, $5) AND type = $6", acNdc, acPharmacist, acYear, acMonth, acDay, acType)
		LogSql(fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s", "SELECT qty, type, date, script, " +
			"id FROM orderdb WHERE ndc = '", acNdc, "' AND pharmacist = '", acPharmacist, "' AND date = make_date(",
			acYear, ", ", acMonth, ", ", acDay, ") AND type = '", acType, "';"))
	} else if acType == "Actual Count" {
		rows, err = db.Query("SELECT qty, type, date, script, id FROM orderdb WHERE ndc = $1 AND pharmacist = $2 AND " +
			"date = make_date($3, $4, $5) AND type = $6 AND script = $7", acNdc, acPharmacist, acYear, acMonth, acDay,
			acType, acScript)
		LogSql(fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s", "SELECT qty, type, date, script, " +
			"id FROM orderdb WHERE ndc = '", acNdc, "' AND pharmacist = '", acPharmacist, "' AND date = make_date(",
			acYear, ", ", acMonth, ", ", acDay, ") AND type = '", acType, " AND script = '", acScript, "';"))
	} else {
		rows, err = db.Query("SELECT qty, type, date, script, id FROM orderdb WHERE ndc = $1 AND pharmacist = $2 AND " +
			"date = make_date($3, $4, $5) AND script = $6", acNdc, acPharmacist, acYear, acMonth, acDay, acScript)
		LogSql(fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s", "SELECT qty, type, date, script, " +
			"id FROM orderdb WHERE ndc = '", acNdc, "' AND pharmacist = '", acPharmacist, "' AND date = make_date(",
			acYear, ", ", acMonth, ", ", acDay, ")  AND script = '", acScript, "';"))
	}

	issue(err)

	for rows.Next() {
		err := rows.Scan(&lnQty, &lcType, &lcDate, &acScript, &lnId)
		issue(err)
		lasOrder = append(lasOrder, MakeOrder(acNdc, acPharmacist, acScript, lcType, lnQty, lcDate, lnId))
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	return lasOrder
}

/**
Function: DeleteOrder
Description: Delete the order with the given id.
@param acId The id to DELETE
 */
func DeleteOrder(anId int64) {
	var lnQty float64
	var lcNdc string
	var lcType string
	rows, err := db.Query("SELECT qty, ndc, type FROM orderdb WHERE id = $1", anId)
	issue(err)
	LogSql(fmt.Sprintf("%s%d%s", "SELECT qty, ndc FROM orderdb WHERE id = ", anId, ";"))

	rows.Next()
	err = rows.Scan(&lnQty, &lcNdc, &lcType)
	issue(err)

	if lcType == "Purchase" {
		_, err = db.Query("UPDATE drugdb set qty = qty - $1 WHERE ndc = $2", lnQty, lcNdc)
		issue(err)
		LogSql(fmt.Sprintf("%s%f%s%s%s", "UPDATE drugdb SET qty = qty - ", lnQty, " WHERE ndc = '", lcNdc, "';"))

		_, err = db.Query("UPDATE orderdb set qty = qty - $1 WHERE ndc = $2 AND id > $3 AND type = 'Actual Count'",
			lnQty, lcNdc, anId)
		issue(err)
		LogSql(fmt.Sprintf("%s%f%s%s%s%d%s", "UPDATE orderdb set qty = qty - ", lnQty, " WHERE ndc = '", lcNdc,
			"', AND id > ", anId, " AND type = 'Actual Count';"))
	} else if lcType == "Prescription" {
		_, err = db.Query("UPDATE drugdb set qty = qty + $1 WHERE ndc = $2", lnQty, lcNdc)
		issue(err)
		LogSql(fmt.Sprintf("%s%f%s%s%s", "UPDATE drugdb SET qty = qty + ", lnQty, " WHERE ndc = '", lcNdc, "';"))

		_, err = db.Query("UPDATE orderdb set qty = qty + $1 WHERE ndc = $2 AND id > $3 AND type = 'Actual Count'",
			lnQty, lcNdc, anId)
		issue(err)
		LogSql(fmt.Sprintf("%s%f%s%s%s%d%s", "UPDATE orderdb set qty = qty + ", lnQty, " WHERE ndc = '", lcNdc,
			"', AND id > ", anId, " AND type = 'Actual Count';"))
	}

	if lcType != "Audit" && lcType != "Actual Count" {
	}

	_, err = db.Query("DELETE FROM orderdb WHERE id = $1", anId)
	issue(err)
	LogSql(fmt.Sprintf("%s%d%s", "DELETE FROM orderdb WHERE id =", anId, ";"))

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}

/**
Function: UpdateOrder
Description: Updates the quantity of an order specified by the passed in id. This
also update all Actual Counts after the specified orders AND the total drug qty
@param acId The id of the order to edit
@param acScript The script number of the order
@param acQty The new quantity
 */
func UpdateOrder(acId string, acScript string, acQty string) {
	var err error
	var rows *sql.Rows
	var lnOldQty float64
	var lnOldDrugDBQty float64
	var lcNdc string
	lnQty, _ := strconv.ParseFloat(acQty, 64)
	lnId, err := strconv.ParseInt(acId, 10, 64)

	rows, err = db.Query("SELECT qty, ndc FROM orderdb WHERE id = $1 ", lnId)
	issue(err)
	LogSql(fmt.Sprintf("%s%d%s", "SELECT qty, ndc FROM orderdb WHERE id = ", lnId, ";"))

	rows.Next()
	err = rows.Scan(&lnOldQty, &lcNdc)
	issue(err)

	_, err = db.Query("UPDATE orderdb SET qty = $1, script = $2 WHERE id = $3", lnQty, acScript, lnId)
	issue(err)
	LogSql(fmt.Sprintf("%s%f%s%s%s%d%s", "UPDATE orderdb SET qty = ", lnQty, ", script = '", acScript,
		"' WHERE id = ", lnId, ";"))

	lnDifference := lnQty - lnOldQty

	// Fix the drugDB value as well
	rows, err = db.Query("SELECT qty FROM drugdb WHERE ndc = $1", lcNdc)
	issue(err)
	LogSql( fmt.Sprintf("%s%s%s", "SELECT qty FROM drugdb WHERE ndc = '", lcNdc, "';"))

	rows.Next()
	err = rows.Scan(&lnOldDrugDBQty)

	_, err = db.Query("UPDATE drugdb SET qty = $1 WHERE ndc = $2", lnOldDrugDBQty + lnDifference, lcNdc)
	issue(err)
	LogSql(fmt.Sprintf("%s%f%s%s%s", "UPDATE drugdb SET qty = ", lnOldDrugDBQty + lnDifference, " WHERE ndc ='",
		lcNdc, "';"))

	// Update the other orders that came after the UPDATEd order
	rows, err = db.Query("SELECT qty, id FROM orderdb WHERE id > $1 AND type = 'Actual Count' AND ndc = $2", lnId, lcNdc)
	issue(err)
	LogSql(fmt.Sprintf( "%s%d%s%s%s", "SELECT qty, id FROM orderdb WHERE id > ", lnId,
		" AND type = 'Actual Count' AND ndc = '", lcNdc, ";"))

	var lnActualCount   float64
	var lnActualCountId int64
	for rows.Next() {
		err = rows.Scan(&lnActualCount, &lnActualCountId)
		issue(err)

		_, err = db.Query("UPDATE orderdb SET qty = $1 WHERE id = $2", lnActualCount + lnDifference, lnActualCountId)
		issue(err)
		LogSql(fmt.Sprintf("%s%f%s%d%s", "UPDATE orderdb SET qty = ", lnActualCount + lnDifference, " WHERE id = ",
			lnActualCountId, ";"))
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()
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
	lnQty, _ := strconv.ParseFloat(anQty, 64)

	row, err := db.Query("SELECT count(script) FROM orderdb WHERE script = $1 AND "+
		"date = make_date($2, $3, $4) AND qty = $5 AND ndc = $6 AND type = $7;", acScript, lnYear, lnMonth, lnDay, lnQty,
		acNdc, acOrderType)

	issue(err)

	LogSql(fmt.Sprintf("%s%s%s%d%s%d%s%d%s%f%s%s%s%s%s", "SELECT count(script) FROM orderdb WHERE script = '",
		acScript, "' AND date = make_date(", lnYear, ", ", lnMonth, ", ", lnDay, ") AND qty = ", lnQty, " AND ndc = '",
		acNdc, "', AND type = '", acOrderType, "';"))

	lnCount := 0
	row.Next()
	err = row.Scan(&lnCount)
	issue(err)

	if lnCount != 0 {
		return false
	}

	_, err = db.Query("INSERT INTO orderdb (ndc, pharmacist, qty, date, logdate, script, type) "+
		"VALUES ($1, $2, $3, make_date($4, $5, $6), current_date, $7, $8);", acNdc, acPharmacist, lnQty, lnYear, lnMonth,
		lnDay, acScript, acOrderType)
	issue(err)
	LogSql(fmt.Sprintf("%s%s%s%s%s%f%s%d%s%d%s%d%s%s%s%s%s", "INSERT INTO orderdb (ndc, pharmacist, qty, date, " +
		"logdate, script, type) VALUES ('", acNdc, "', '", acPharmacist, "', ", lnQty, ", make_date(", lnYear, ", ",
		lnMonth, ", ", lnDay, "), current_date, '", acScript, "', ", acOrderType, ");"))

	defer func() {
		if err := row.Close(); err != nil {
			log.Fatal(err)
		}
	}()

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

	lnQty, err := strconv.ParseFloat(acQty, 64)
	issue(err)
	rows, err := db.Query("SELECT qty FROM drugdb WHERE ndc = $1", acNdc)
	issue(err)
	LogSql( fmt.Sprintf("%s%s%s", "SELECT qty FROM drugdb WHERE ndc = '", acNdc, "';"))
	rows.Next()
	err = rows.Scan(&lnRowQty)
	if err != nil {
		_, err = db.Query("INSERT INTO drugdb (ndc, qty) VALUES ($1, $2)", acNdc, lnQty)
		issue(err)
		LogSql(fmt.Sprintf("%s%s%s%f%s", "INSERT INTO drugdb (ndc, qty) VALUES ('", acNdc, "', ", lnQty, ");"))
		return
	}
	lnNewQty := lnRowQty - lnQty
	_, err = db.Query("UPDATE drugdb SET qty = $1 WHERE ndc = $2", lnNewQty, acNdc)
	issue(err)
	LogSql(fmt.Sprintf("%s%f%s%s%s", "UPDATE drugdb SET qty = ", lnNewQty, " WHERE ndc ='", acNdc, "';"))

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()
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
	rows, err := db.Query("SELECT qty FROM drugdb WHERE ndc = $1", acNdc)
	issue(err)
	LogSql( fmt.Sprintf("%s%s%s", "SELECT qty FROM drugdb WHERE ndc = '", acNdc, "';"))
	rows.Next()
	err = rows.Scan(&lnRowQty)
	lnQty, err := strconv.Atoi(acQty)
	issue(err)
	_, err = db.Query("UPDATE drugdb SET qty = $1 WHERE ndc = $2", lnQty, acNdc)
	issue(err)
	LogSql( fmt.Sprintf("%s%d%s%s%s", "UPDATE drugdb SET qty = ", lnQty, " WHERE ndc ='", acNdc, "';"))

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()
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
		lcQtyDiff := setDrugQty(acNdc, acActual)
		lbCheck = addType(acNdc, acPharmacist, acMonth, acDay, acYear,
			strconv.Itoa(lcQtyDiff), acScript, "Actual Count")
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

	alterQty(acNdc, "-" + acQty)

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
	row, err := db.Query("SELECT count(ndc) FROM drugdb WHERE ndc = $1", acNdc)
	issue(err)
	LogSql( fmt.Sprintf("%s%s%s", "SELECT count(ndc) FROM drugdb WHERE ndc = '", acNdc, "';"))
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
	_, err = db.Query("INSERT INTO drugdb (ndc, date) VALUES ($1, make_date($2, $3, $4))",
		acNdc, lnYear, lnMonth, lnDay)
	issue(err)
	LogSql(fmt.Sprintf("%s%s%s%d%s%d%s%d%s", "INSERT INTO drugdb (ndc, date) VALUES ('", acNdc, "', make_date(",
		lnYear,	", ", lnMonth, ", ", lnDay, "));"))
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
	_, err := db.Query("UPDATE drugdb SET size = $1, form = $2, item_num = $3, name = $4 WHERE ndc = $5",
		acSize, acForm, acItemNum, acName, acNdc)
	issue(err)
	LogSql(fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s", "UPDATE drugdb SET size = '", acSize, "', form = '", acForm, "', " +
		"item_num = '", acItemNum, "', name = '", acName, "', WHERE ndc = '", acNdc, "';"))
}
