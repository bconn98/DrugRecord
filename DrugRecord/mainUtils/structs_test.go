package mainUtils

/**
File: structs_test
Description: Tests all the functions in structs
@author Bryan Conn
@date 6/23/18
 */

import (
	"testing"
)

/**
Function: TestMakePurchase
Description: Checks if a purchase correctly adds to the drug quantity
 */
func TestMakePurchase(t *testing.T) {
	testDrug1 := Drug{"Test", "999-9999-999", 110}
	testDrug2 := Drug{"Test", "999-9999-999", 100}
	testDate := Date{6, 23, 18}
	testPurch := Purchase{testDrug1, testDate, 10}
	makePurch := MakePurchase(testDrug2, testDate, 10)
	if testPurch != makePurch {
		t.Error("The purchases don't match!")
	}
}

/**
Function: TestMakeAudit
Description: Checks if the audit values match and if they don't match
 */
func TestMakeAudit(t *testing.T) {
	testDrug := Drug{"Test", "999-9999-999", 110}
	testADate := Date{6, 23, 18}
	testLDate := Date{6, 23, 18}
	testAudit1 := Audit{testDrug, "BRC", 110, testADate, testLDate}
	makeAudit := MakeAudit(testDrug, 110, "BRC", testADate, testLDate)
	if testAudit1 != makeAudit {
		t.Error("The audits don't match!")
	}
	testAudit2 := Audit{testDrug, "BRC", 100, testADate, testLDate}
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
	testDate1 := Date{6, 23, 18}
	testDate2 := MakeDate("6", "23", "18")
	if testDate1 != testDate2 {
		t.Error("The dates don't match!")
	}
}

/**
Function: TestMakeOrder
Description: Makes sure that an order doesn't change any information in a purchase
 */
func TestMakeOrder(t *testing.T) {
	testDrug1 := Drug{"Test", "999-9999-999", 100}
	testDate := Date{6, 23, 18}
	testPurch := Purchase{testDrug1, testDate, 10}
	testOrder := Order{testPurch}
	makeOrder := MakeOrder(testPurch)
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
	testDrug1 := Drug{"Test", "999-9999-999", 60}
	testDrug2 := Drug{"Test", "999-9999-999", 110}
	testDate := Date{6, 23, 18}
	testLDate := Date{6, 23, 18}
	testPrescrip := Prescription{testDrug1, "BRC", "999", 50, testDate, testLDate}
	makePrescrip := MakePrescription(testDrug2, testDate, 50, "BRC", "999", testLDate)
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