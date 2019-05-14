package mainUtils

import (
	"time"
	"strconv"
)

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
	Day int
	Month time.Month
	Year int
}

/**
Function: MakeDate
Description: Makes a drug struct with the month, day, year
@param month The month
@param day The day
@param year The year
@return A Date object
 */
 func MakeDate(month time.Month, day int, year int) Date {
 	return Date{day, month, year}
 }

/**
Prescription struct contains the ndc of  drug of the order, the pharmacist that filled
the order, the script id, the quantity of the order, the date the order was filled
 */
type Prescription struct {
	ndc string
	Pharmacist, Script string
	OrderQuantity float64
	date time.Time
}

/**
Function: MakePrescription
Description: Makes a Prescription struct
@param ndc The ndc of the drug being ordered
@param qty The quantity of the order
@param pharm The initials of the pharmacist
@param script The script id
@param date The date of the order
@return A prescription object
 */
func MakePrescription(ndc string, pharm string, script string, qty float64, date time.Time) Prescription {
	return Prescription{ndc, pharm, script, qty, date}
}

/**
Audit struct contains an audited quantity, the pharmacist who performed the audit,
the date the audit was performed and the ndc of the auditted drug
 */
type Audit struct {
	ndc string
	Pharmacist string
	AuditQuantity float64
	date time.Time
}

/**
Function: MakeAudit
Description: Makes an audit struct
@param qty The quantity recorded in the audit
@param pharm The initials of the pharmacist who performed the audit
@param date The the audit was performed
@param ndc The ndc of the drug
@return An Audit object
 */
func MakeAudit(ndc string, pharm string, qty float64, date time.Time) Audit {
	return Audit{ndc, pharm, qty, date}
}

/**
Purchase struct contains the ndc of a drug, purchase date, and purchased quantity,
and the pharmacist that counted the drug
 */
type Purchase struct {
	ndc string
	pharm string
	Qty float64
	date time.Time
}

/**
Function: MakePurchase
Description: Makes a Purchase struct
@param ndc The ndc of the drug that was bought
@param date The date the purchase was added to the supply
@param qty The quantity bought
@param pharm The pharmacist that counted the drug
@return A Purchase object
 */
func MakePurchase(ndc string, pharm string, qty float64, date time.Time) Purchase {
	return Purchase{ndc, pharm, qty, date}
}

/**
Drug struct contains an id name, ndc code, and quantity
 */
type Drug struct {
	Id, NDC string
	Quantity float64
}

/**
Function: makeDrug
Description: Given: a drug name, ndc, and quantity, creates a drug structure
@param name The name of the drug
@param ndc The ndc specific to the drug
@param qty The current quantity of the drug
 */
func MakeDrug(name string, ndc string, qty float64) Drug {
	return Drug{name, ndc,qty}
}

/**
Function: UpdateQty
Description: Updates the quantity of the drug
 */
 func (drug Drug) UpdateQty(qty float64) Drug {
 	qty = drug.Quantity + qty
 	return MakeDrug(drug.Id, drug.NDC, qty)
 }

/**
Order struct contains the pharmacist on the order, the script/type of the order
the quantity and the date of the order
 */
type Order struct {
	Pharm string
	Script, Typ string
	Qty float64
	Date string
}

/**
Function: MakeOrder
Description: Creates an order using an audit, prescription, or purchase
@param pharm The pharmacist on the order
@param script The script/type of the order
@param qty The quantity of the order
@param date The date of the order
@return An Order Object
 */
func MakeOrder(pharm string, script string, typ string, qty float64, date time.Time) Order {
	var dateS string
	dateS = date.Month().String() + " " + strconv.Itoa(date.Day()) + " " + strconv.Itoa(date.Year())
	return Order{pharm, script, typ, qty, dateS}
}

/**
User struct contains a username and a password value
 */
type User struct {
	UserName string
	PassVal int
}