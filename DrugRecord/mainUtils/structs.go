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

func ParseDateStrings(acTime time.Time) (string, string, string) {
	lcMonth := acTime.Month().String()
	lcDay := strconv.Itoa(acTime.Day())
	lcYear := strconv.Itoa(acTime.Year())
	return lcMonth, lcDay, lcYear
}

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
	McNdc                  string
	mcPharmacist, mcScript string
	mcYear                 string
	mcMonth                string
	mcDay                  string
	mnOrderQuantity        float64
	mrActualQty            float64
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
func MakePrescription(acNdc string, asPharmacist string, acScript string, anQty float64, acYear string, acMonth string,
	acDay string, arActualQty float64) Prescription {
	return Prescription{acNdc, asPharmacist, acScript, acYear, acMonth,
		acDay, anQty, arActualQty}
}

/**
Audit struct contains an audited quantity, the pharmacist who performed the audit,
the date the audit was performed and the ndc of the audited drug
*/
type Audit struct {
	mcNdc           string
	mcPharmacist    string
	mcYear          string
	mcMonth         string
	mcDay           string
	mnAuditQuantity float64
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
func MakeAudit(acNdc string, acPharmacist string, anAuditQuantity float64, acYear string, acMonth string,
	acDay string) Audit {
	return Audit{acNdc, acPharmacist, acYear, acMonth,
		acDay, anAuditQuantity}
}

/**
Purchase struct contains the ndc of a drug, purchase date, and purchased quantity,
and the pharmacist that counted the drug
*/
type Purchase struct {
	MnNdc        string
	mcPharmacist string
	mcInvoice    string
	mcYear       string
	mcMonth      string
	mcDay        string
	mrQty        float64
	mrActualQty  float64
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
func MakePurchase(acNdc string, acPharmacist string, acInvoice string, acYear string, acMonth string, acDay string,
	anQty float64, anActualQty float64) Purchase {
	return Purchase{acNdc, acPharmacist, acInvoice, acYear, acMonth,
		acDay, anQty, anActualQty}
}

/**
Drug struct contains an id name, ndc code, and quantity
*/
type Drug struct {
	McNdc      string
	MrQuantity float64
	McName     string
	McDate     time.Time
	McForm     string
	McSize     string
	McItemNum  string
}

/**
Function: makeDrug
Description: Given: a drug name, ndc, and quantity, creates a drug structure
@param acName The name of the drug
@param acNdc The ndc specific to the drug
@param anQty The current quantity of the drug
*/
// func MakeDrug(acName string, acNdc string, anQty float64) Drug {
// 	return Drug{acName, acNdc, anQty}
// }

/**
Function: UpdateQty
Description: Updates the quantity of the drug
@param anQty The quantity to change by
*/
// func (drug Drug) UpdateQty(anQty float64) Drug {
// 	anQty = drug.MnQuantity + anQty
// 	return MakeDrug(drug.McId, drug.McNdc, anQty)
// }

/**
Order struct contains the pharmacist on the order, the script/type of the order
the quantity and the date of the order
*/
type Order struct {
	AcPharmacist       string
	AcMonth            string
	AcDay              int
	AcYear             int
	AcScript, AcType   string
	ArQty, ArActualQty float64
	AcNdc              string
	AnId               int64
}

/**
Function: MakeOrder
Description: Creates an order using an audit, prescription, or purchase
@param acPharmacist The pharmacist on the order
@param acScript The script/type of the order
@param anQty The quantity of the order
@param anQty The real quantity of the order drug
@param acDate The date of the order
@param acType The type of the order
@return An Order Object
*/
func MakeOrder(acNdc string, acPharmacist string, acScript string, acType string, arQty float64, arActualQty float64,
	acYear string, acMonth string, acDay string, anId int64) Order {

	lnDay, _ := strconv.Atoi(acDay)
	lnYear, _ := strconv.Atoi(acYear)

	return Order{acPharmacist, acMonth, lnDay, lnYear,
		acScript, acType, arQty, arActualQty, acNdc, anId}
}

/**
User struct contains a username and a password value
*/
type User struct {
	UserName string
	PassVal  int
}

type NewDrug struct {
	Error string
	Ndc   string
	Id    int
}
