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

	selectString := fmt.Sprintf("%s%s%s",
		"SELECT pharmacist, date, qty, script, type, id FROM orderdb WHERE ndc = '", acNdc, "' ORDER BY date desc;")

	rows, err := db.Query("SELECT pharmacist, date, qty, script, type, "+
		"id FROM orderdb WHERE ndc = $1 ORDER BY date desc;", acNdc)
	issue(err)
	Log(selectString, SQL)

	for rows.Next() {
		if rows.Err() != nil {
			issue(rows.Err())
			break
		}

		issue(rows.Scan(&lcPharmacist, &lcDate, &lnQty, &lcScript, &lcType, &lnId))
		lcMonth, lcDay, lcYear := ParseDateStrings(lcDate)
		lasOrders = append(lasOrders, MakeOrder(acNdc, lcPharmacist, lcScript, lcType, lnQty, 0, lcYear,
			lcMonth, lcDay, lnId))
	}

	selectString = fmt.Sprintf("%s%s%s",
		"SELECT name, form, item_num, size, date, qty FROM drugdb WHERE ndc = '", acNdc, "';")

	err = db.QueryRow("SELECT name, form, item_num, size, date, qty FROM drugdb WHERE ndc = $1;",
		acNdc).Scan(&lcName, &lcForm, &lcItemNum, &lcSize, &lcDate, &lnDrugQty)
	issue(err)
	Log(selectString, SQL)

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
 * @param order The order id to get
 */
func GetOrder(anId int64) []Order {

	var err error
	var lcNdc, lcPharm, lcType, lcScript string
	var lcDate time.Time
	var lnQty float64
	var rows *sql.Rows
	var lasOrders []Order
	var selectString string

	rows, err = db.Query("SELECT ndc, pharmacist, qty, date, type, script FROM orderdb WHERE id = $1", anId)

	selectString = fmt.Sprintf("%s%d%s", "SELECT ndc, pharmacist, qty, date, "+
		"type FROM orderdb WHERE id = ", anId, ";")

	Log(selectString, SQL)
	issue(err)

	for rows.Next() {
		if rows.Err() != nil {
			issue(rows.Err())
			break
		}

		issue(rows.Scan(&lcNdc, &lcPharm, &lnQty, &lcDate, &lcType, &lcScript))
		lasOrders = append(lasOrders, MakeOrder(lcNdc, lcPharm, lcScript, lcType, lnQty,
			0, strconv.Itoa(lcDate.Year()), lcDate.Month().String(), strconv.Itoa(lcDate.Day()), anId))
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
	Log(selectString, SQL)

	if strings.ToUpper(lcType) != "AUDIT" {
		if strings.ToUpper(lcType) == "PURCHASE" {
			lnQty *= -1
		}

		alterQty(lcNdc, -lnQty)

		updateString := fmt.Sprintf("%s%f%s%s%s%d%s", "UPDATE orderdb set qty = qty + ", lnQty,
			" WHERE ndc = '", lcNdc, "' AND id > ", anId, " AND type = 'Over/Short';")

		_, err = db.Exec("UPDATE orderdb set qty = qty + $1 WHERE ndc = $2 AND id > $3 AND type = 'Over/Short';",
			lnQty, lcNdc, anId)
		issue(err)
		Log(updateString, SQL)
	}

	deleteString := fmt.Sprintf("%s%d%s", "DELETE FROM orderdb WHERE id =", anId, ";")
	_, err = db.Exec("DELETE FROM orderdb WHERE id = $1;", anId)
	issue(err)
	Log(deleteString, SQL)

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
	Log(selectString, SQL)

	lcType = strings.ToUpper(lcType)

	updateString := fmt.Sprintf("%s%f%s%s%s%d%s", "UPDATE orderdb SET qty = ", lnQty, ", script = '",
		acScript, "' WHERE id = ", lnId, ";")
	_, err = db.Exec("UPDATE orderdb SET qty = $1, script = $2 WHERE id = $3;", lnQty, acScript, lnId)
	issue(err)
	Log(updateString, SQL)

	var lrDifference float64
	if lcType == "PURCHASE" || lcType == "OVER/SHORT" {
		lrDifference = lnQty - lnOldQty
	} else {
		lrDifference = lnOldQty - lnQty
	}

	if lcType != "AUDIT" && lcType != "OVER/SHORT" {
		// Fix the drugDB value as well
		alterQty(lcNdc, -lrDifference)
	}
}

/**
 * Function: addType
 * Description: Creates a new row in the orderdb with passed in attributes
 * @param order The order to be added to the DB
 */
func addType(order Order) (bool, int) {

	lnId := 0
	var lnCount int

	// Check and see if the order is already entered in the system
	selectString := fmt.Sprintf("%s%s%s%d%s%s%s%d%s%f%s%f%s%s%s", "SELECT count(script) FROM orderdb WHERE script = '",
		order.AcScript, "' AND date = make_date(", order.AcYear, ", ", order.AcMonth, ", ", order.AcDay, ") AND qty = ",
		order.ArQty, " AND ndc = '", order.ArQty, "' AND type = '", order.AcType, "';")

	issue(db.QueryRow("SELECT count(script) FROM orderdb WHERE script = $1 AND date = make_date($2, $3, "+
		"$4) AND qty = $5 AND ndc = $6 AND type = $7;", order.AcScript, order.AcYear, order.AcMonth, order.AcDay,
		order.ArQty, order.AcNdc, order.AcType).Scan(&lnCount))
	Log(selectString, SQL)

	if lnCount != 0 {
		return false, 0
	}

	insertString := fmt.Sprintf("%s%s%s%s%s%f%s%d%s%s%s%d%s%s%s%s%s", "INSERT INTO orderdb (ndc, pharmacist, qty, "+
		"date, logdate, script, type) VALUES ('", order.AcNdc, "', '", order.AcPharmacist, "', ", order.ArQty, ", make_date(",
		order.AcYear, ", ", order.AcMonth, ", ", order.AcDay, "), current_date, '", order.AcScript, "', ", order.AcType, ");")

	_, err = db.Exec("INSERT INTO orderdb (ndc, pharmacist, qty, date, "+
		"logdate, script, type) VALUES ($1, $2, $3, make_date($4, $5, $6), current_date, $7, $8);", order.AcNdc,
		order.AcPharmacist, order.ArQty, order.AcYear, order.AcMonth, order.AcDay, order.AcScript, order.AcType)

	issue(err)
	Log(insertString, SQL)

	issue(db.QueryRow("SELECT id FROM orderdb where ndc = $1 and pharmacist = $2 and qty = $3 and date = make_date("+
		"$4, $5, $6) and script = $7 and type = $8;", order.AcNdc, order.AcPharmacist, order.ArQty, order.AcYear,
		order.AcMonth, order.AcDay, order.AcScript, order.AcType).Scan(&lnId))

	return true, lnId
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
	Log(selectString, SQL)

	if err != nil {
		insertString := fmt.Sprintf("%s%s%s%f%s", "INSERT INTO drugdb (ndc, qty) VALUES ('", acNdc, "', ", arQty, ");")
		_, err = db.Exec("INSERT INTO drugdb (ndc, qty) VALUES ($1, $2);", acNdc, arQty)
		issue(err)
		Log(insertString, SQL)
		return
	}

	lnNewQty := lnRowQty - arQty
	updateString := fmt.Sprintf("%s%.3f%s%s%s", "UPDATE drugdb SET qty = ", lnNewQty, " WHERE ndc =", acNdc, ";")
	_, err = db.Exec("UPDATE drugdb SET qty = $1 WHERE ndc = $2", lnNewQty, acNdc)
	issue(err)
	Log(updateString, SQL)
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
	Log(selectString, SQL)

	updateString := fmt.Sprintf("%s%f%s%s%s", "UPDATE drugdb SET qty = ", arQty, " WHERE ndc ='", acNdc, "';")
	_, err = db.Exec("UPDATE drugdb SET qty = $1 WHERE ndc = $2;", arQty, acNdc)
	issue(err)
	Log(updateString, SQL)

	return lnRowQty - arQty
}

/**
 * Function: AddPrescription
 * Description: Adds a prescription type order to the orderdb
 * @param prescription The prescription to add to the DB
 */
func AddPrescription(prescription Prescription) (bool, int) {

	lbCheck := true
	order := MakeOrder(prescription.McNdc, prescription.mcPharmacist, prescription.mcScript, "Prescription",
		prescription.mnOrderQuantity, prescription.mrActualQty, prescription.mcYear, prescription.mcMonth,
		prescription.mcDay, 0) // Id is not important

	check, id := addType(order)
	if !check {
		return false, id
	}

	alterQty(prescription.McNdc, prescription.mnOrderQuantity)

	if prescription.mrActualQty != -1000 {
		lnQtyDiff := setDrugQty(order.AcNdc, prescription.mrActualQty)
		order.AcType = "Over/Short"
		order.ArQty = lnQtyDiff
		lbCheck, _ = addType(order)
	}
	return lbCheck, id
}

/**
 * Function: AddAudit
 * Description: Adds a audit type order to the orderdb
 * @param audit The audit to add to the DB
 */
func AddAudit(audit Audit) (bool, int) {

	issue(err)

	lbCheck := true
	order := MakeOrder(audit.mcNdc, audit.mcPharmacist, "", "Audit", audit.mnAuditQuantity,
		audit.mnAuditQuantity, audit.mcYear, audit.mcMonth, audit.mcDay, 0) // Id is not important

	check, id := addType(order)
	if !check {
		return false, id
	}

	lnQtyDiff := setDrugQty(order.AcNdc, audit.mnAuditQuantity)
	order.AcType = "Over/Short"
	order.ArQty = lnQtyDiff
	lbCheck, _ = addType(order)

	return lbCheck, id
}

/**
 * Function: AddPurchase
 * Description: Adds a purchase type order to the orderdb
 * @param purchase The purchase to add to the DB
 */
func AddPurchase(purchase Purchase) (bool, int) {

	lbCheck := true
	order := MakeOrder(purchase.MnNdc, purchase.mcPharmacist, purchase.mcInvoice, "Purchase", purchase.mrQty,
		purchase.mrActualQty, purchase.mcYear, purchase.mcMonth, purchase.mcDay, 0) // Id is not important

	check, id := addType(order)
	if !check {
		return false, id
	}

	alterQty(purchase.MnNdc, -1*purchase.mrQty)

	if purchase.mrActualQty != -1000 {
		lnQtyDiff := setDrugQty(purchase.MnNdc, purchase.mrActualQty)
		order.AcType = "Over/Short"
		order.ArQty = lnQtyDiff
		lbCheck, _ = addType(order)
	}
	return lbCheck, id
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
	Log(selectString, SQL)

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
	Log(insertString, SQL)
}

/**
 * Function: UpdateDrug
 * Description: Adds the correct default values to a drug
 * @param acSize The size of the packet
 * @param acForm The form of the drug
 * @param acItemNum The item number of the drug
 * @param acName The name of the drug
 * @param acNdc The ndc of the drug
 * @param acOldNdc The ndc of the drug on first entry
 */
func UpdateDrug(acSize string, acForm string, acItemNum string, acName string, acNdc string,
	arQty float64, acOldNdc string) {
	updateString := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s", "UPDATE drugdb SET size = '", acSize, "', form = '",
		acForm, "', item_num = '", acItemNum, "', name = '", acName, "', ndc = '", acNdc, "' WHERE ndc = '", acOldNdc, "';")

	_, err := db.Exec("UPDATE drugdb SET size = $1, form = $2, item_num = $3, name = $4, ndc = $5, "+
		"qty = $6 WHERE ndc = $7;",
		acSize, acForm, acItemNum, acName, acNdc, arQty, acOldNdc)
	issue(err)
	Log(updateString, SQL)

	if acNdc != acOldNdc {
		_, err = db.Exec("DELETE from drugdb where ndc = $1", acOldNdc)
	}
}

/**
 * Method: UpdateOrderNdc
 * Description: Reset the ndc of an order
 * @param acId The id of the order to change
 * @param acNdc The new ndc
 */
func UpdateOrderNdc(acId string, acNdc string) {
	var lrQty float64
	var lcNdc string
	var lcType string

	issue(db.QueryRow("SELECT ndc FROM orderdb WHERE id = $1", acId).Scan(&lcNdc))

	if lcNdc != acNdc {
		_, err := db.Exec("DELETE FROM drugdb WHERE ndc = $1", lcNdc)
		issue(err)
	}

	updateString := fmt.Sprintf("%s%s%s%s%s", "UPDATE orderdb set ndc = '", acNdc, "' where id = '", acId, "';")

	_, err = db.Exec("UPDATE orderdb set ndc = $1 where id = $2", acNdc, acId)
	issue(err)
	Log(updateString, SQL)

	issue(db.QueryRow("SELECT qty, type FROM orderdb WHERE id = $1;", acId).Scan(&lrQty, &lcType))

	if lcType == "Purchase" {
		lrQty *= -1
	}
	updateString = fmt.Sprintf("%s%f%s%s%s", "UPDATE drugdb set qty = qty - ", lrQty, " where ndc = '", acNdc, "';")
	_, err = db.Exec("UPDATE drugdb set qty = qty - $1 where ndc = $2", lrQty, acNdc)
	issue(err)
	Log(updateString, SQL)
}

func GetDrugs(acName string) []DrugDB {
	var ndc, name, size, form, item_num, qty string
	var date time.Time
	var lasDrugs []DrugDB
	acName = "%" + acName + "%"

	rows, err := db.Query("SELECT ndc, name, size, form, item_num, qty, "+
		"date from drugdb where lower(name) like lower($1)", acName)
	issue(err)

	for rows.Next() {
		if rows.Err() != nil {
			issue(rows.Err())
			break
		}

		issue(rows.Scan(&ndc, &name, &size, &form, &item_num, &qty, &date))

		lrQty, err := strconv.ParseFloat(qty, 10)
		issue(err)
		month, day, year := ParseDateStrings(date)

		lasDrugs = append(lasDrugs,
			DrugDB{
				Name:     name,
				Ndc:      ndc,
				Size:     size,
				Form:     form,
				ItemNum:  item_num,
				Quantity: lrQty,
				Month:    month,
				Day:      day,
				Year:     year,
			})
	}

	return lasDrugs

}

func GetDrug(acNdc string) Drug {
	var name, size, form, item_num, qty string
	issue(db.QueryRow("SELECT name, size, form, item_num, qty from drugdb where ndc = $1", acNdc).Scan(&name, &size,
		&form, &item_num, &qty))
	lrQty, err := strconv.ParseFloat(qty, 10)
	issue(err)

	return Drug{
		McNdc:      acNdc,
		MrQuantity: lrQty,
		McName:     name,
		McDate:     time.Time{},
		McForm:     form,
		McSize:     size,
		McItemNum:  item_num,
	}
}
