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
 * Function: FindNDC
 * Description: Given a NDC, finds AND returns all corresponding rows
 * @param ndc The NDC to match
 * @return An array of orders with the given NDC, vital drug information
 */
func FindNDC(acNdc string) (string, string, string, string, string, time.Time, float64, []Order) {
	var lnId int64
	var lcDate time.Time
	var lasOrders []Order
	var lnQty, lnDrugQty float64
	var lcPharmacist, lcScript, lcType, lcName, lcForm, lcItemNum, lcSize string

	selectString := fmt.Sprintf("%s%s%s", "SELECT pharmacist, date, qty, script, type, id FROM orderdb WHERE ndc = '",
		acNdc, "' ORDER BY date desc;")
	rows, err := db.Query(selectString)
	issue(err)
	LogSql(selectString)

	for rows.Next() {
		if rows.Err() != nil {
			issue(rows.Err())
			break
		}

		issue(rows.Scan(&lcPharmacist, &lcDate, &lnQty, &lcScript, &lcType, &lnId))
		lasOrders = append(lasOrders, MakeOrder(acNdc, lcPharmacist, lcScript, lcType, lnQty, lcDate, lnId))
	}

	selectString = fmt.Sprintf("%s%s%s", "SELECT name, form, item_num, size, date, qty FROM drugdb WHERE ndc = '", acNdc, "';")
	err = db.QueryRow(selectString).Scan(&lcName, &lcForm, &lcItemNum, &lcSize, &lcDate, &lnDrugQty)
	issue(err)
	LogSql(selectString)

	if err != nil {
		return "", "", "", "", "", time.Time{}, 0, nil
	}

	defer func() {
		issue(rows.Close())
	}()

	return lcName, acNdc, lcForm, lcItemNum, lcSize, lcDate, lnDrugQty, lasOrders
}

/**
 * Function: GetOrder
 * Description: Gets the fields of an order that weren't specified by user
 * @param acNdc The ndc of the drug to get
 * @param acPharmacist The pharmacist who logged the order
 * @param acMonth The month the order was logged
 * @param acDay The day the order was logged
 * @param acYear The year the order was logged
 * @param acScript The script number of the order
 * @param acType The type of the order
 */
func GetOrder(acNdc string, acPharmacist string, acMonth string, acDay string, acYear string,
	acScript string, acType string) []Order {

	var err error
	var lnId int64
	var lnQty float64
	var rows *sql.Rows
	var lasOrder []Order
	var lcDate time.Time
	var lcType, selectString string

	if acType == "Audit" {
		selectString = fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s", "SELECT qty, type, date, script, "+
			"id FROM orderdb WHERE ndc = '", acNdc, "' AND pharmacist = '", acPharmacist, "' AND date = make_date(",
			acYear, ", ", acMonth, ", ", acDay, ") AND type = '", acType, "';")

	} else if acType == "Actual Count" {
		selectString = fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s", "SELECT qty, type, date, script, "+
			"id FROM orderdb WHERE ndc = '", acNdc, "' AND pharmacist = '", acPharmacist, "' AND date = make_date(",
			acYear, ", ", acMonth, ", ", acDay, ") AND type = '", acType, " AND script = '", acScript, "';")
	} else {
		selectString = fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s", "SELECT qty, type, date, script, "+
			"id FROM orderdb WHERE ndc = '", acNdc, "' AND pharmacist = '", acPharmacist, "' AND date = make_date(",
			acYear, ", ", acMonth, ", ", acDay, ")  AND script = '", acScript, "';")
	}

	rows, err = db.Query(selectString)
	LogSql(selectString)
	issue(err)

	for rows.Next() {
		if rows.Err() != nil {
			issue(rows.Err())
			break
		}

		issue(rows.Scan(&lnQty, &lcType, &lcDate, &acScript, &lnId))
		lasOrder = append(lasOrder, MakeOrder(acNdc, acPharmacist, acScript, lcType, lnQty, lcDate, lnId))
	}

	defer func() {
		issue(rows.Close())
	}()

	return lasOrder
}

/**
 * Function: DeleteOrder
 * Description: Delete the order with the given id.
 * @param acId The id to DELETE
 */
func DeleteOrder(anId int64) {
	var lnQty float64
	var lcNdc, lcType string

	selectString := fmt.Sprintf("%s%d%s", "SELECT qty, ndc FROM orderdb WHERE id = ", anId, ";")
	issue(db.QueryRow(selectString).Scan(&lnQty, &lcNdc, &lcType))
	LogSql(selectString)

	if lcType == "Purchase" {
		updateString := fmt.Sprintf("%s%f%s%s%s", "UPDATE drugdb SET qty = qty - ", lnQty, " WHERE ndc = '", lcNdc, "';")

		_, err = db.Exec(updateString)
		issue(err)
		LogSql(updateString)

		updateString = fmt.Sprintf("%s%f%s%s%s%d%s", "UPDATE orderdb set qty = qty - ", lnQty, " WHERE ndc = '",
			lcNdc, "', AND id > ", anId, " AND type = 'Actual Count';")

		_, err = db.Exec(updateString)
		issue(err)
		LogSql(updateString)

	} else if lcType == "Prescription" {
		updateString := fmt.Sprintf("%s%f%s%s%s", "UPDATE drugdb SET qty = qty + ", lnQty, " WHERE ndc = '", lcNdc, "';")

		_, err = db.Exec(updateString)
		issue(err)
		LogSql(updateString)

		updateString = fmt.Sprintf("%s%f%s%s%s%d%s", "UPDATE orderdb set qty = qty + ", lnQty, " WHERE ndc = '",
			lcNdc, "', AND id > ", anId, " AND type = 'Actual Count';")

		_, err = db.Exec(updateString)
		issue(err)
		LogSql(updateString)
	}

	deleteString := fmt.Sprintf("%s%d%s", "DELETE FROM orderdb WHERE id =", anId, ";")
	_, err = db.Exec(deleteString)
	issue(err)
	LogSql(deleteString)

}

/**
 * Function: UpdateOrder
 * Description: Updates the quantity of an order specified by the passed in id. This
 * also update all Actual Counts after the specified orders AND the total drug qty
 * @param acId The id of the order to edit
 * @param acScript The script number of the order
 * @param acQty The new quantity
 */
func UpdateOrder(acId string, acScript string, acQty string) {
	var err error
	var lcNdc string
	var rows *sql.Rows
	var lnActualCountId int64
	var lnOldQty, lnOldDrugDBQty, lnActualCount float64

	lnQty, _ := strconv.ParseFloat(acQty, 64)
	lnId, err := strconv.ParseInt(acId, 10, 64)

	selectString := fmt.Sprintf("%s%d%s", "SELECT qty, ndc FROM orderdb WHERE id = ", lnId, ";")
	issue(db.QueryRow(selectString).Scan(&lnOldQty, &lcNdc))
	LogSql(selectString)

	updateString := fmt.Sprintf("%s%f%s%s%s%d%s", "UPDATE orderdb SET qty = ", lnQty, ", script = '", acScript,
		"' WHERE id = ", lnId, ";")
	_, err = db.Exec(updateString)
	issue(err)
	LogSql(updateString)

	lnDifference := lnQty - lnOldQty

	// Fix the drugDB value as well
	selectString = fmt.Sprintf("%s%s%s", "SELECT qty FROM drugdb WHERE ndc = '", lcNdc, "';")
	issue(db.QueryRow(selectString).Scan(&lnOldDrugDBQty))
	LogSql(selectString)

	updateString = fmt.Sprintf("%s%f%s%s%s", "UPDATE drugdb SET qty = ", lnOldDrugDBQty+lnDifference,
		" WHERE ndc ='", lcNdc, "';")
	_, err = db.Exec(updateString)
	issue(err)
	LogSql(updateString)

	// Update the other orders that came after the UPDATEd order
	selectString = fmt.Sprintf("%s%d%s%s%s", "SELECT qty, id FROM orderdb WHERE id > ", lnId,
		" AND type = 'Actual Count' AND ndc = '", lcNdc, ";")
	rows, err = db.Query(selectString)
	issue(err)
	LogSql(selectString)

	for rows.Next() {

		if rows.Err() != nil {
			issue(rows.Err())
			break
		}

		err = rows.Scan(&lnActualCount, &lnActualCountId)
		issue(err)

		updateString = fmt.Sprintf("%s%f%s%d%s", "UPDATE orderdb SET qty = ", lnActualCount+lnDifference,
			" WHERE id = ", lnActualCountId, ";")
		_, err = db.Exec(updateString)
		issue(err)
		LogSql(updateString)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}

/**
 * Function: addType
 * Description: Creates a new row in the orderdb with passed in attributes
 * @param acNdc The NDC of the drug of interest
 * @param acPharmacist The pharmacist who is inputting the order
 * @param acMonth The month the order was made
 * @param anDay The day the order was made
 * @param anYear The year the order was made
 * @param anQty The quantity of the order
 * @param acScript The order#/Script# or blank
 * @param acOrderType The type of the order
 */
func addType(acNdc string, acPharmacist string, acMonth string, anDay string, anYear string,
	anQty string, acScript string, acOrderType string) bool {

	var lnCount int

	lnMonth, _ := strconv.Atoi(acMonth)
	lnDay, _ := strconv.Atoi(anDay)
	lnYear, _ := strconv.Atoi(anYear)
	lnQty, _ := strconv.ParseFloat(anQty, 64)

	selectString := fmt.Sprintf("%s%s%s%d%s%d%s%d%s%f%s%s%s%s%s", "SELECT count(script) FROM orderdb WHERE script = '",
		acScript, "' AND date = make_date(", lnYear, ", ", lnMonth, ", ", lnDay, ") AND qty = ", lnQty, " AND ndc = '",
		acNdc, "', AND type = '", acOrderType, "';")

	issue(db.QueryRow(selectString).Scan(&lnCount))
	LogSql(selectString)

	if lnCount != 0 {
		return false
	}

	insertString := fmt.Sprintf("%s%s%s%s%s%f%s%d%s%d%s%d%s%s%s%s%s", "INSERT INTO orderdb (ndc, pharmacist, qty, date, "+
		"logdate, script, type) VALUES ('", acNdc, "', '", acPharmacist, "', ", lnQty, ", make_date(", lnYear, ", ",
		lnMonth, ", ", lnDay, "), current_date, '", acScript, "', ", acOrderType, ");")
	_, err = db.Exec(insertString)
	issue(err)
	LogSql(insertString)

	return true
}

/**
 * Function: alterQty
 * Description: Alters a drugs quantity using its NDC to find it
 * @param acNdc The ndc value of the drug in question
 * @param acQty The quantity of the alteration
 */
func alterQty(acNdc string, acQty string) {
	var lnRowQty float64

	if strings.Contains(acQty, "--") {
		acQty = strings.Replace(acQty, "--", "-", 1)
	}

	lnQty, err := strconv.ParseFloat(acQty, 64)
	issue(err)

	selectString := fmt.Sprintf("%s%s%s", "SELECT qty FROM drugdb WHERE ndc = '", acNdc, "';")
	err = db.QueryRow(selectString).Scan(&lnRowQty)
	issue(err)
	LogSql(selectString)

	if err != nil {
		insertString := fmt.Sprintf("%s%s%s%f%s", "INSERT INTO drugdb (ndc, qty) VALUES ('", acNdc, "', ", lnQty, ");")
		_, err = db.Exec(insertString)
		issue(err)
		LogSql(insertString)
		return
	}

	lnNewQty := lnRowQty - lnQty
	updateString := fmt.Sprintf("%s%f%s%s%s", "UPDATE drugdb SET qty = ", lnNewQty, " WHERE ndc ='", acNdc, "';")
	_, err = db.Exec(updateString)
	issue(err)
	LogSql(updateString)
}

/**
 * Function: setDrugQty
 * Description: Sets the drug quantity to the new value
 * @param acNdc The ndc value of the drug in question
 * @param acQty The new qty of the drug
 * @return The difference in change
 */
func setDrugQty(acNdc string, acQty string) int {
	var lnRowQty int

	selectString := fmt.Sprintf("%s%s%s", "SELECT qty FROM drugdb WHERE ndc = '", acNdc, "';")
	issue(db.QueryRow(selectString).Scan(&lnRowQty))
	LogSql(selectString)

	lnQty, err := strconv.Atoi(acQty)
	issue(err)

	updateString := fmt.Sprintf("%s%d%s%s%s", "UPDATE drugdb SET qty = ", lnQty, " WHERE ndc ='", acNdc, "';")
	_, err = db.Exec(updateString)
	issue(err)
	LogSql(updateString)

	return lnRowQty - lnQty
}

/**
 * Function: AddPrescription
 * Description: Adds a prescription type order to the orderdb
 * @param acNdc The NDC of the drug of interest
 * @param acPharmacist The pharmacist who is inputting the prescription
 * @param acMonth The month the prescription was made
 * @param acDay The day the prescription was made
 * @param acYear The year the prescription was made
 * @param acQty The quantity of the prescription
 * @param acScript The prescription number
 * @param acActual The actual count if entered
 */
func AddPrescription(acNdc string, acPharmacist string, acMonth string, acDay string,
	acYear string, acQty string, acScript string, acActual string) bool {

	var lbCheck bool

	if !addType(acNdc, acPharmacist, acMonth, acDay, acYear, acQty, acScript, "Prescription") {
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
 * Function: AddAudit
 * Description: Adds a audit type order to the orderdb
 * @param acNdc The NDC of the drug of interest
 * @param acPharmacist The pharmacist who is inputting the audit
 * @param acMonth The month the audit was made
 * @param acDay The day the audit was made
 * @param acYear The year the audit was made
 * @param acQty The quantity of the audit
 * @param acActual The actual count if entered
 */
func AddAudit(acNdc string, acPharmacist string, acMonth string, acDay string,
	acYear string, acQty string, acActual string) bool {

	var lbCheck bool

	if !addType(acNdc, acPharmacist, acMonth, acDay, acYear, acQty, "", "Audit") {
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
 * Function: AddPurchase
 * Description: Adds a purchase type order to the orderdb
 * @param acNdc The NDC of the drug of interest
 * @param acPharmacist The pharmacist who is inputting the purchase
 * @param acMonth The month the purchase was made
 * @param acDay The day the purchase was made
 * @param acYear The year the purchase was made
 * @param acQty The quantity of the purchase
 * @param acActual The actual count if entered
 */
func AddPurchase(acNdc string, acPharmacist string, acMonth string, acDay string,
	acYear string, acQty string, order string, acActual string) bool {

	var lbCheck bool

	if !addType(acNdc, acPharmacist, acMonth, acDay, acYear, acQty, order, "Purchase") {
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
 * Function: NewCheck
 * Description: See if the drug is in the database yet
 * @param acNdc The NDC of the drug to check
 * @return If the drug is in the database
 */
func NewCheck(acNdc string) bool {
	var lnCount int

	selectString := fmt.Sprintf("%s%s%s", "SELECT count(ndc) FROM drugdb WHERE ndc = '", acNdc, "';")
	issue(db.QueryRow(selectString).Scan(&lnCount))
	LogSql(selectString)

	if lnCount < 1 {
		return false
	} else {
		return true
	}
}

/**
 * Function: AddDrug
 * Description: Adds the drug to the database without the defaults
 * @param acNdc The ndc
 * @param acMonth The month as a string
 * @param acDay The day as a string
 * @param acYear The year as a string
 */
func AddDrug(acNdc string, acMonth string, acDay string, acYear string) {
	lnMonth, _ := strconv.Atoi(acMonth)
	lnDay, _ := strconv.Atoi(acDay)
	lnYear, _ := strconv.Atoi(acYear)

	insertString := fmt.Sprintf("%s%s%s%d%s%d%s%d%s", "INSERT INTO drugdb (ndc, date) VALUES ('", acNdc,
		"', make_date(", lnYear, ", ", lnMonth, ", ", lnDay, "));")

	_, err = db.Exec(insertString)
	issue(err)
	LogSql(insertString)
}

/**
 * Function: UpdateDrug
 * Description: Adds the correct default values to a drug
 * @param acSize The size of the packet
 * @param acForm The form of the drug
 * @param acItemNum The item number of the drug
 * @param acName The name of the drug
 * @param acNdc The ndc of the drug
 */
func UpdateDrug(acSize string, acForm string, acItemNum string, acName string, acNdc string) {
	updateString := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s", "UPDATE drugdb SET size = '", acSize, "', form = '",
		acForm, "', "+"item_num = '", acItemNum, "', name = '", acName, "', WHERE ndc = '", acNdc, "';")
	_, err := db.Exec(updateString)
	issue(err)
	LogSql(updateString)
}
