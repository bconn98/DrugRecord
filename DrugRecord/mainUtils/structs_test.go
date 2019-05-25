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
	var ndc = "99999-9999-99"
	var pharm = "BRC"
	var qty = 10.0
	d := time.Date(2018, 6, 23, 12, 30, 0, 0, time.UTC)
	testPurch := Purchase{ndc, pharm, qty, d}
	makePurch := MakePurchase(ndc, pharm, qty, d)
	if testPurch != makePurch {
		t.Error("The purchases don't match!")
	}
}

/**
Function: TestMakeAudit
Description: Checks if the audit values match and if they don't match
*/
func TestMakeAudit(t *testing.T) {
	var ndc = "99999-9999-99"
	var pharm = "BRC"
	var qty = 110.0
	var qty2 = 100.0
	d := time.Date(2018, 6, 23, 12, 30, 0, 0, time.UTC)
	testAudit1 := Audit{ndc, pharm, qty, d}
	makeAudit := MakeAudit(ndc, pharm, qty, d)
	if testAudit1 != makeAudit {
		t.Error("The audits don't match!")
	}
	testAudit2 := Audit{ndc, pharm, qty2, d}
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
	d := time.Date(2018, 6, 23, 12, 30, 0, 0, time.UTC)
	testDate1 := Date{23, 6, 2018}
	testDate2 := MakeDate(d.Month(), d.Day(), d.Year())
	if testDate1 != testDate2 {
		t.Error("The dates don't match!")
	}
}

/**
Function: TestMakeOrder
Description: Makes sure that an order doesn't change any information in a purchase
*/
func TestMakeOrder(t *testing.T) {
	var ndc = "99999-9999-99"
	var pharm = "BRC"
	var typ = "PURCHASE"
	var qty = 100.0
	d := time.Date(2018, 6, 23, 12, 30, 0, 0, time.UTC)

	testOrder := Order{ndc, pharm, typ, qty, "June 23 2018"}
	makeOrder := MakeOrder(ndc, pharm, typ, qty, d)

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
	var ndc = "999-9999-999"
	var pharm = "BRC"
	var script = "999"
	var qty = 50.0
	d := time.Date(2018, 6, 23, 12, 30, 0, 0, time.UTC)
	testPrescrip := Prescription{ndc, pharm, script, qty, d}
	makePrescrip := MakePrescription(ndc, pharm, script, qty, d)
	if testPrescrip != makePrescrip {
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
