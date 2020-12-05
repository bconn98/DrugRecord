package utils

/**
File: structs_test
Description: Tests all the functions in structs
@author Bryan Conn
@date 6/23/18
*/

import (
	"strconv"
	"testing"
	"time"
)

/**
Function: TestMakePurchase
Description: Checks if a purchase correctly adds to the drug quantity
*/
func TestMakePurchase(t *testing.T) {
	var lcNdc = "99999-9999-99"
	var lcPharmacist = "BRC"
	var lnQty = 10.0
	var lcInvoice = "98034324"
	lcDate := time.Date(2018, 6, 23, 12, 30, 0, 0, time.UTC)
	testPurchase := Purchase{lcNdc, lcPharmacist, lcInvoice, strconv.Itoa(lcDate.Year()),
		lcDate.Month().String(), strconv.Itoa(lcDate.Day()), lnQty, lnQty}
	makePurchase := MakePurchase(lcNdc, lcPharmacist, lcInvoice, strconv.Itoa(lcDate.Year()), lcDate.Month().String(),
		strconv.Itoa(lcDate.Day()), lnQty, lnQty)
	if testPurchase != makePurchase {
		t.Error("The purchases don't match!")
	}
}

/**
Function: TestMakeAudit
Description: Checks if the audit values match and if they don't match
*/
func TestMakeAudit(t *testing.T) {
	var lcNdc = "99999-9999-99"
	var lcPharmacist = "BRC"
	var lrQty = 110.0
	var lrQty2 = 100.0
	lcDate := time.Date(2018, 6, 23, 12, 30, 0, 0, time.UTC)
	testAudit1 := Audit{lcNdc, lcPharmacist, strconv.Itoa(lcDate.Year()), lcDate.Month().String(),
		strconv.Itoa(lcDate.Day()), lrQty}
	makeAudit := MakeAudit(lcNdc, lcPharmacist, lrQty, strconv.Itoa(lcDate.Year()), lcDate.Month().String(),
		strconv.Itoa(lcDate.Day()))
	if testAudit1 != makeAudit {
		t.Error("The audits don't match!")
	}
	testAudit2 := Audit{lcNdc, lcPharmacist, strconv.Itoa(lcDate.Year()), lcDate.Month().String(),
		strconv.Itoa(lcDate.Day()), lrQty2}
	if testAudit2 == makeAudit {
		t.Error("The audits shouldn't match!")
	}
}

/**
Function: TestMakeDate
Description: Checks if the make date function correctly changes strings into integers
*/
func TestMakeDate(t *testing.T) {
	lcDate := time.Date(2018, 6, 23, 12, 30, 0, 0, time.UTC)
	testDate1 := Date{23, 6, 2018}
	testDate2 := MakeDate(lcDate.Month(), lcDate.Day(), lcDate.Year())
	if testDate1 != testDate2 {
		t.Error("The dates don't match!")
	}
}

/**
Function: TestMakeOrder
Description: Makes sure that an order doesn't change any information in a purchase
*/
func TestMakeOrder(t *testing.T) {
	var lcNdc = "99999-9999-99"
	var lcScript = "99999-9999-99"
	var lcPharmacist = "BRC"
	var lcType = "PURCHASE"
	var lrQty = 100.0
	var lnId int64
	lnId = 1
	lcDate := time.Date(2018, 6, 23, 12, 30, 0, 0, time.UTC)

	testOrder := Order{lcPharmacist, lcDate.Month().String(), lcDate.Day(), lcDate.Year(), lcScript, lcType, lrQty,
		lrQty, lcNdc, lnId}
	makeOrder := MakeOrder(lcNdc, lcPharmacist, lcScript, lcType, lrQty, lrQty, strconv.Itoa(lcDate.Year()),
		lcDate.Month().String(), strconv.Itoa(lcDate.Day()), lnId)

	if testOrder != makeOrder {
		t.Error("The orders don't match!")
	}
}

/**
Function: TestMakePrescription
Description: Makes sure that the make function correctly decreases the quantity
*/
func TestMakePrescription(t *testing.T) {
	// Value is lower because I'm not calling the updateQty
	var lcNdc = "999-9999-999"
	var lcPharmacist = "BRC"
	var script = "999"
	var lrQty = 50.0
	var lcDate = time.Date(2018, 6, 23, 12, 30, 0, 0, time.UTC)

	testPrescription := Prescription{lcNdc, lcPharmacist, script, strconv.Itoa(lcDate.Year()),
		lcDate.Month().String(), strconv.Itoa(lcDate.Day()), lrQty, lrQty}
	makePrescription := MakePrescription(lcNdc, lcPharmacist, script, lrQty, strconv.Itoa(lcDate.Year()),
		lcDate.Month().String(), strconv.Itoa(lcDate.Day()), lrQty)
	if testPrescription != makePrescription {
		t.Error("The prescriptions don't match!")
	}
}
