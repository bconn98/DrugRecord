package goDB

import "strconv"

/**
File: structs.go
Description: All the structs needed to implement the C2 record
@author Bryan Conn
@date 5/16/18
 */

/**
Date struct holds a month day and year
 */
type Date struct {
	Month, Day, Year int
}

/**
Function: MakeDate
Description: Makes a drug struct with the month, day, year
@param month The month
@param day The day
@param year The year
 */
 func MakeDate(monthS string, dayS string, yearS string) Date {
	 month, _ := strconv.Atoi(monthS)
	 day, _ := strconv.Atoi(dayS)
	 year, _ := strconv.Atoi(yearS)
 	return Date{month, day, year}
 }

/**
Prescription struct contains the drug of the order, the pharmacist that filled the order,
the script id, the quantity of the order, the date the order was filled, the date the
order was logged
 */
type Prescription struct {
	OrderDrug Drug
	Pharmacist, Script string
	OrderQuantity int
	OrderDate, LogDate Date
}

/**
Function: MakePrescription
Description: Makes a Prescription struct
@param drug The drug being ordered
@param oDate The date the order was filled
@param oQty The quantity of the order
@param pharm The initials of the pharmacist
@param script The script id
@param lDate The date the order was logged
 */
func MakePrescription(drug Drug, oDate Date, oQty int, pharm string, script string, lDate Date) Prescription {
	drug = drug.UpdateQty(-oQty)
	return Prescription{drug, pharm, script, oQty,
		oDate, lDate}
}

/**
Audit struct contains an audited quantity, the pharmacist who performed the audit,
the date the audit was performed and the date the audit was logged
 */
type Audit struct {
	ADrug Drug
	Pharmacist string
	AuditQuantity int
	AuditDate, LogDate Date
}

/**
Function: MakeAudit
Description: Makes an audit struct
@param qty The quantity recorded in the audit
@param pharm The initials of the pharmacist who performed the audit
@param oDate The the audit was performed
@param lDate The date logged
 */
func MakeAudit(drug Drug, qty int, pharm string, oDate Date, lDate Date) Audit {
	return Audit{drug, pharm, qty, oDate, lDate}
}

/**
Purchase struct contains a drug, purchase date, and purchased quantity
 */
type Purchase struct {
	PurchasedDrug Drug
	PurchaseDate Date
	Qty int
}

/**
Function: MakePurchase
Description: Makes a Purchase struct
@param drug The drug that was bought
@param date The date the purchase was added to the supply
@param qty The quantity bought
 */
func MakePurchase(drug Drug, date Date, qty int) Purchase {
	drug = drug.UpdateQty(qty)
	return Purchase{drug, date, qty}
}

/**
Drug struct contains an id name, ndc code, and quantity
 */
type Drug struct {
	Id, NDC string
	Quantity int
}

/**
Function: makeDrug
Description: Given: a drug name, ndc, and quantity, creates a drug structure
@param name The name of the drug
@param ndc The ndc specific to the drug
@param qty The current quantity of the drug
 */
func MakeDrug(name string, ndc string, qty int) Drug {
	return Drug{name, ndc,qty}
}

/**
Function: UpdateQty
Description: Updates the quantity of the drug
 */
 func (drug Drug) UpdateQty(qty int) Drug {
 	qty = drug.Quantity + qty
 	return MakeDrug(drug.Id, drug.NDC, qty)
 }

/**
Order struct contains either an audit, prescription, or purchase
 */
type Order struct {
	ThisOrder interface{}
}

/**
Function: MakeOrder
Description: Creates an order using an audit, prescription, or purchase
@param thisOrder Either an audit, prescription, or purchase struct
 */
func MakeOrder(thisOrder interface{}) Order {
	return Order{thisOrder}
}