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
	"strconv"
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

	selectString := fmt.Sprintf("%s%s%s",
		"SELECT pharmacist, date, qty, script, type, id FROM orderdb WHERE ndc = '", acNdc, "' ORDER BY date desc;")

	rows, err := db.Query("SELECT pharmacist, date, qty, script, type, "+
		"id FROM orderdb WHERE ndc = $1 ORDER BY date desc;", acNdc)
	issue(err)
	LogSql(selectString)

	for rows.Next() {
		if rows.Err() != nil {
			issue(rows.Err())
			break
		}

		issue(rows.Scan(&lcPharmacist, &lcDate, &lnQty, &lcScript, &lcType, &lnId))
		lcMonth, lcDay, lcYear := ParseDateStrings(lcDate)
		lasOrders = append(lasOrders, MakeOrder(acNdc, lcPharmacist, lcScript, lcType, lnQty, lcYear, lcMonth, lcDay,
			lnId))
	}

	selectString = fmt.Sprintf("%s%s%s",
		"SELECT name, form, item_num, size, date, qty FROM drugdb WHERE ndc = '", acNdc, "';")

	err = db.QueryRow("SELECT name, form, item_num, size, date, qty FROM drugdb WHERE ndc = $1;",
		acNdc).Scan(&lcName, &lcForm, &lcItemNum, &lcSize, &lcDate, &lnDrugQty)
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
 * @param order The order to get from the
 */
func GetOrder(order Order) []Order {

	var err error
	var lnId int64
	var lnQty float64
	var rows *sql.Rows
	var lasOrders []Order
	var selectString string

	if order.AcType == "Audit" {
		selectString = fmt.Sprintf("%s%s%s%s%s%d%s%s%s%d%s%s%s", "SELECT qty, id FROM orderdb WHERE ndc = '",
			order.AcNdc, "' AND pharmacist = '", order.AcPharmacist, "' AND date = make_date(",
			order.AcYear, ", ", order.AcMonth, ", ", order.AcDay, ") AND type = '", order.AcType, "';")

		rows, err = db.Query("SELECT qty, id FROM orderdb WHERE ndc = $1 AND pharmacist = $2 AND "+
			"date = make_date($3, $4, $5) AND type = $6;", order.AcNdc, order.AcPharmacist, order.AcYear,
			order.AcMonth, order.AcDay, order.AcType)

	} else {
		selectString = fmt.Sprintf("%s%s%s%s%s%d%s%s%s%d%s%s%s%s%s", "SELECT qty, id FROM orderdb WHERE ndc = '",
			order.AcNdc, "' AND pharmacist = '", order.AcPharmacist, "' AND date = make_date(", order.AcYear, ", ",
			order.AcMonth, ", ", order.AcDay, ") AND type = '", order.AcType, " AND script = '", order.AcScript, "';")

		rows, err = db.Query("SELECT qty, id FROM orderdb WHERE ndc = $1 AND pharmacist = $2 AND date = make_date($3, "+
			"$4, "+
			"$5) AND type = $6 AND script = $7;", order.AcNdc, order.AcPharmacist, order.AcYear, order.AcMonth, order.AcDay,
			order.AcType, order.AcScript)
	}

	LogSql(selectString)
	issue(err)

	for rows.Next() {
		if rows.Err() != nil {
			issue(rows.Err())
			break
		}

		issue(rows.Scan(&lnQty, &lnId))
		lasOrders = append(lasOrders, MakeOrder(order.AcNdc, order.AcPharmacist, order.AcScript, order.AcType, lnQty,
			strconv.Itoa(order.AcYear), order.AcMonth, strconv.Itoa(order.AcDay), lnId))
	}

	defer func() {
		issue(rows.Close())
	}()

	return lasOrders
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
	issue(db.QueryRow("SELECT qty, ndc, type FROM orderdb WHERE id = $1;", anId).Scan(&lnQty, &lcNdc, &lcType))
	LogSql(selectString)

	if lcType != "Audit" {
		if lcType == "Purchase" {
			lnQty *= -1
		}

		alterQty(lcNdc, -lnQty)

		updateString := fmt.Sprintf("%s%f%s%s%s%d%s", "UPDATE orderdb set qty = qty + ", lnQty,
			" WHERE ndc = '", lcNdc, "' AND id > ", anId, " AND type = 'Over/Short';")

		_, err = db.Exec("UPDATE orderdb set qty = qty + $1 WHERE ndc = $2 AND id > $3 AND type = 'Over/Short';",
			lnQty, lcNdc, anId)
		issue(err)
		LogSql(updateString)
	}

	deleteString := fmt.Sprintf("%s%d%s", "DELETE FROM orderdb WHERE id =", anId, ";")
	_, err = db.Exec("DELETE FROM orderdb WHERE id = $1;", anId)
	issue(err)
	LogSql(deleteString)

}

/**
 * Function: UpdateOrder
 * Description: Updates the quantity of an order specified by the passed in id. This
 * also update all Over/Shorts after the specified orders AND the total drug qty
 * @param acId The id of the order to edit
 * @param acScript The script number of the order
 * @param acQty The new quantity
 */
func UpdateOrder(acId string, acScript string, acQty string) {
	var err error
	var lcNdc, lcType string
	var lnOldQty float64

	lnQty, _ := strconv.ParseFloat(acQty, 64)
	lnId, err := strconv.ParseInt(acId, 10, 64)

	selectString := fmt.Sprintf("%s%d%s", "SELECT qty, ndc FROM orderdb WHERE id = ", lnId, ";")
	issue(db.QueryRow("SELECT qty, ndc, type FROM orderdb WHERE id = $1;", lnId).Scan(&lnOldQty, &lcNdc, &lcType))
	LogSql(selectString)

	updateString := fmt.Sprintf("%s%f%s%s%s%d%s", "UPDATE orderdb SET qty = ", lnQty, ", script = '",
		acScript, "' WHERE id = ", lnId, ";")
	_, err = db.Exec("UPDATE orderdb SET qty = $1, script = $2 WHERE id = $3;", lnQty, acScript, lnId)
	issue(err)
	LogSql(updateString)

	var lrDifference float64
	if lcType == "Purchase" || lcType == "Over/Short" {
		lrDifference = lnQty - lnOldQty
	} else {
		lrDifference = lnOldQty - lnQty
	}

	if lcType != "Audit" && lcType != "Over/Short" {
		// Fix the drugDB value as well
		alterQty(lcNdc, -lrDifference)
	}
}

/**
 * Function: addType
 * Description: Creates a new row in the orderdb with passed in attributes
 * @param order The order to be added to the DB
 */
func addType(order Order) bool {

	var lnCount int

	selectString := fmt.Sprintf("%s%s%s%d%s%s%s%d%s%f%s%f%s%s%s", "SELECT count(script) FROM orderdb WHERE script = '",
		order.AcScript, "' AND date = make_date(", order.AcYear, ", ", order.AcMonth, ", ", order.AcDay, ") AND qty = ",
		order.ArQty, " AND ndc = '", order.ArQty, "' AND type = '", order.AcType, "';")

	issue(db.QueryRow("SELECT count(script) FROM orderdb WHERE script = $1 AND date = make_date($2, $3, "+
		"$4) AND qty = $5 AND ndc = $6 AND type = $7;", order.AcScript, order.AcYear, order.AcMonth, order.AcDay,
		order.ArQty, order.AcNdc, order.AcType).Scan(&lnCount))
	LogSql(selectString)

	if lnCount != 0 {
		return false
	}

	insertString := fmt.Sprintf("%s%s%s%s%s%f%s%d%s%s%s%d%s%s%s%s%s", "INSERT INTO orderdb (ndc, pharmacist, qty, "+
		"date, logdate, script, type) VALUES ('", order.AcNdc, "', '", order.AcPharmacist, "', ", order.ArQty, ", make_date(",
		order.AcYear, ", ", order.AcMonth, ", ", order.AcDay, "), current_date, '", order.AcScript, "', ", order.AcType, ");")

	_, err = db.Exec("INSERT INTO orderdb (ndc, pharmacist, qty, date, "+
		"logdate, script, type) VALUES ($1, $2, $3, make_date($4, $5, $6), current_date, $7, $8);", order.AcNdc,
		order.AcPharmacist, order.ArQty, order.AcYear, order.AcMonth, order.AcDay, order.AcScript, order.AcType)

	issue(err)
	LogSql(insertString)

	return true
}

/**
 * Function: alterQty
 * Description: Alters a drugs quantity using its NDC to find it
 * @param acNdc The ndc value of the drug in question
 * @param arQty The quantity of the alteration
 */
func alterQty(acNdc string, arQty float64) {
	var lnRowQty float64

	selectString := fmt.Sprintf("%s%s%s", "SELECT qty FROM drugdb WHERE ndc = '", acNdc, "';")
	err = db.QueryRow("SELECT qty FROM drugdb WHERE ndc = $1;", acNdc).Scan(&lnRowQty)
	issue(err)
	LogSql(selectString)

	if err != nil {
		insertString := fmt.Sprintf("%s%s%s%f%s", "INSERT INTO drugdb (ndc, qty) VALUES ('", acNdc, "', ", arQty, ");")
		_, err = db.Exec("INSERT INTO drugdb (ndc, qty) VALUES ($1, $2);", acNdc, arQty)
		issue(err)
		LogSql(insertString)
		return
	}

	lnNewQty := lnRowQty - arQty
	updateString := fmt.Sprintf("%s%.3f%s%s%s", "UPDATE drugdb SET qty = ", lnNewQty, " WHERE ndc =", acNdc, ";")
	_, err = db.Exec("UPDATE drugdb SET qty = $1 WHERE ndc = $2", lnNewQty, acNdc)
	issue(err)
	LogSql(updateString)
}

/**
 * Function: setDrugQty
 * Description: Sets the drug quantity to the new value
 * @param acNdc The ndc value of the drug in question
 * @param arQty The new qty of the drug
 * @return The difference in change
 */
func setDrugQty(acNdc string, arQty float64) float64 {
	var lnRowQty float64

	selectString := fmt.Sprintf("%s%s%s", "SELECT qty FROM drugdb WHERE ndc = '", acNdc, "';")
	issue(db.QueryRow("SELECT qty FROM drugdb WHERE ndc = $1;", acNdc).Scan(&lnRowQty))
	LogSql(selectString)

	updateString := fmt.Sprintf("%s%f%s%s%s", "UPDATE drugdb SET qty = ", arQty, " WHERE ndc ='", acNdc, "';")
	_, err = db.Exec("UPDATE drugdb SET qty = $1 WHERE ndc = $2;", arQty, acNdc)
	issue(err)
	LogSql(updateString)

	return lnRowQty - arQty
}

/**
 * Function: AddPrescription
 * Description: Adds a prescription type order to the orderdb
 * @param prescription The prescription to add to the DB
 */
func AddPrescription(prescription Prescription) bool {

	lbCheck := true
	order := MakeOrder(prescription.mcNdc, prescription.mcPharmacist, prescription.mcScript, "Prescription",
		prescription.mnOrderQuantity, prescription.mcYear, prescription.mcMonth, prescription.mcDay,
		0) // Id is not important

	if !addType(order) {
		return false
	}

	alterQty(prescription.mcNdc, prescription.mnOrderQuantity)

	if prescription.mrActualQty != -1000 {
		lnQtyDiff := setDrugQty(order.AcNdc, prescription.mrActualQty)
		order.AcType = "Over/Short"
		order.ArQty = lnQtyDiff
		lbCheck = addType(order)
	}
	return lbCheck
}

/**
 * Function: AddAudit
 * Description: Adds a audit type order to the orderdb
 * @param audit The audit to add to the DB
 */
func AddAudit(audit Audit) bool {

	issue(err)

	lbCheck := true
	order := MakeOrder(audit.mcNdc, audit.mcPharmacist, "", "Audit", audit.mnAuditQuantity,
		audit.mcYear, audit.mcMonth, audit.mcDay, 0) // Id is not important

	if !addType(order) {
		return false
	}

	lnQtyDiff := setDrugQty(order.AcNdc, audit.mnAuditQuantity)
	order.AcType = "Over/Short"
	order.ArQty = lnQtyDiff
	lbCheck = addType(order)

	return lbCheck
}

/**
 * Function: AddPurchase
 * Description: Adds a purchase type order to the orderdb
 * @param purchase The purchase to add to the DB
 */
func AddPurchase(purchase Purchase) bool {

	lbCheck := true
	order := MakeOrder(purchase.mnNdc, purchase.mcPharmacist, purchase.mcInvoice, "Purchase", purchase.mrQty,
		purchase.mcYear, purchase.mcMonth, purchase.mcDay, 0) // Id is not important

	if !addType(order) {
		return false
	}

	alterQty(purchase.mnNdc, -1*purchase.mrQty)

	if purchase.mrActualQty != -1000 {
		lnQtyDiff := setDrugQty(purchase.mnNdc, purchase.mrActualQty)
		order.AcType = "Over/Short"
		order.ArQty = lnQtyDiff
		lbCheck = addType(order)
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
	issue(db.QueryRow("SELECT count(ndc) FROM drugdb WHERE ndc = $1;", acNdc).Scan(&lnCount))
	LogSql(selectString)

	return lnCount >= 1
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

	_, err = db.Exec("INSERT INTO drugdb (ndc, date) VALUES ($1, make_date($2, $3, $4));", acNdc, lnYear, lnMonth, lnDay)
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
		acForm, "', item_num = '", acItemNum, "', name = '", acName, "' WHERE ndc = '", acNdc, "';")

	_, err := db.Exec("UPDATE drugdb SET size = $1, form = $2, item_num = $3, name = $4 WHERE ndc = $5;", acSize,
		acForm, acItemNum, acName, acNdc)
	issue(err)
	LogSql(updateString)
}
