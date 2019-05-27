package mainUtils

/**
File: structs_test
Description: Tests all the functions in structs
@author Bryan Conn
@date 6/23/18
*/

import (
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
	lcDate := time.Date(2018, 6, 23, 12, 30, 0, 0, time.UTC)
	testPurchase := Purchase{lcNdc, lcPharmacist, lnQty, lcDate}
	makePurchase := MakePurchase(lcNdc, lcPharmacist, lnQty, lcDate)
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
	var lnQty = 110.0
	var lnQty2 = 100.0
	lcDate := time.Date(2018, 6, 23, 12, 30, 0, 0, time.UTC)
	testAudit1 := Audit{lcNdc, lcPharmacist, lnQty, lcDate}
	makeAudit := MakeAudit(lcNdc, lcPharmacist, lnQty, lcDate)
	if testAudit1 != makeAudit {
		t.Error("The audits don't match!")
	}
	testAudit2 := Audit{lcNdc, lcPharmacist, lnQty2, lcDate}
	if testAudit2 == makeAudit {
		t.Error("The audits shouldn't match!")
	}
}

/**
Function: TestMakeDrug
Description: Checks if the make drug function correctly returns a drug struct
*/
func TestMakeDrug(t *testing.T) {
	testDrug1 := Drug{"Test", "999-9999-999", 100}
	testDrug2 := MakeDrug("Test", "999-9999-999", 100)
	if testDrug2 != testDrug1 {
		t.Error("The drugs don't match!")
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
	var lnQty = 100.0
	lcDate := time.Date(2018, 6, 23, 12, 30, 0, 0, time.UTC)

	testOrder := Order{lcPharmacist, "June 23 2018", lcScript, lcType, lnQty}
	makeOrder := MakeOrder(lcNdc, lcPharmacist, lcType, lnQty, lcDate)

	if testOrder != makeOrder {
		t.Error("The orders don't match!")
	}
}

/**
Function: TestMakePrescription
Description: Makes sure that the make function correctly decreases the quantity
*/
func TestMakePrescription(t *testing.T) {
	//Value is lower because I'm not calling the updateQty
	var lcNdc = "999-9999-999"
	var lcPharmacist = "BRC"
	var script = "999"
	var lnQty = 50.0
	lcDate := time.Date(2018, 6, 23, 12, 30, 0, 0, time.UTC)
	testPrescription := Prescription{lcNdc, lcPharmacist, script, lnQty, lcDate}
	makePrescription := MakePrescription(lcNdc, lcPharmacist, script, lnQty, lcDate)
	if testPrescription != makePrescription {
		t.Error("The prescriptions don't match!")
	}
}

/**
Function: TestDrug_UpdateQty
Description: Makes sure that the quantity is correctly increased
*/
func TestDrug_UpdateQty(t *testing.T) {
	testDrug1 := Drug{"Test", "999-9999-999", 110}
	testDrug2 := Drug{"Test", "999-9999-999", 100}

	testDrug2 = testDrug2.UpdateQty(10)

	if testDrug1 != testDrug2 {
		t.Error("The drug quantities don't match!")
	}
}
