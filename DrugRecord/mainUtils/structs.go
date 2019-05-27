package mainUtils

import (
	"strconv"
	"time"
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
	MnDay   int
	McMonth time.Month
	MnYear  int
}

/**
Function: MakeDate
Description: Makes a drug struct with the month, day, year
@param acMonth The month
@param acDay The day
@param acYear The year
@return A Date object
*/
func MakeDate(acMonth time.Month, acDay int, acYear int) Date {
	return Date{acDay, acMonth, acYear}
}

/**
Prescription struct contains the ndc of  drug of the order, the pharmacist that filled
the order, the script id, the quantity of the order, the date the order was filled
*/
type Prescription struct {
	mcNdc                string
	mcPharmacist, Script string
	mnOrderQuantity      float64
	mcDate               time.Time
}

/**
Function: MakePrescription
Description: Makes a Prescription struct
@param acNdc The ndc of the drug being ordered
@param anQty The quantity of the order
@param asPharmacist The initials of the pharmacist
@param acScript The script id
@param acDate The date of the order
@return A prescription object
*/
func MakePrescription(acNdc string, asPharmacist string, acScript string, anQty float64, acDate time.Time) Prescription {
	return Prescription{acNdc, asPharmacist, acScript, anQty, acDate}
}

/**
Audit struct contains an audited quantity, the pharmacist who performed the audit,
the date the audit was performed and the ndc of the audited drug
*/
type Audit struct {
	mcNdc           string
	mcPharmacist    string
	mnAuditQuantity float64
	mcDate          time.Time
}

/**
Function: MakeAudit
Description: Makes an audit struct
@param anAuditQuantity The quantity recorded in the audit
@param acPharmacist The initials of the pharmacist who performed the audit
@param acDate The the audit was performed
@param acNdc The ndc of the drug
@return An Audit object
*/
func MakeAudit(acNdc string, acPharmacist string, anAuditQuantity float64, acDate time.Time) Audit {
	return Audit{acNdc, acPharmacist, anAuditQuantity, acDate}
}

/**
Purchase struct contains the ndc of a drug, purchase date, and purchased quantity,
and the pharmacist that counted the drug
*/
type Purchase struct {
	mnNdc        string
	mcPharmacist string
	mnQty        float64
	mcDate       time.Time
}

/**
Function: MakePurchase
Description: Makes a Purchase struct
@param acNdc The ndc of the drug that was bought
@param acDate The date the purchase was added to the supply
@param anQty The quantity bought
@param acPharmacist The pharmacist that counted the drug
@return A Purchase object
*/
func MakePurchase(acNdc string, acPharmacist string, anQty float64, acDate time.Time) Purchase {
	return Purchase{acNdc, acPharmacist, anQty, acDate}
}

/**
Drug struct contains an id name, ndc code, and quantity
*/
type Drug struct {
	McId, McNdc string
	MnQuantity  float64
}

/**
Function: makeDrug
Description: Given: a drug name, ndc, and quantity, creates a drug structure
@param acName The name of the drug
@param acNdc The ndc specific to the drug
@param anQty The current quantity of the drug
*/
func MakeDrug(acName string, acNdc string, anQty float64) Drug {
	return Drug{acName, acNdc, anQty}
}

/**
Function: UpdateQty
Description: Updates the quantity of the drug
@param anQty The quantity to change by
*/
func (drug Drug) UpdateQty(anQty float64) Drug {
	anQty = drug.MnQuantity + anQty
	return MakeDrug(drug.McId, drug.McNdc, anQty)
}

/**
Order struct contains the pharmacist on the order, the script/type of the order
the quantity and the date of the order
*/
type Order struct {
	AcPharmacist, AcDate string
	AcScript, AcType     string
	AnQty                float64
}

/**
Function: MakeOrder
Description: Creates an order using an audit, prescription, or purchase
@param acPharmacist The pharmacist on the order
@param acScript The script/type of the order
@param anQty The quantity of the order
@param acDate The date of the order
@param acType The type of the order
@return An Order Object
*/
func MakeOrder(acPharmacist string, acScript string, acType string, anQty float64, acDate time.Time) Order {
	var lcDate string
	lcDate = acDate.Month().String() + " " + strconv.Itoa(acDate.Day()) + " " + strconv.Itoa(acDate.Year())
	return Order{acPharmacist, lcDate, acScript, acType, anQty}
}

/**
User struct contains a username and a password value
*/
type User struct {
	UserName string
	PassVal  int
}
