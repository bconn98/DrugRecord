package utils

import (
	excelize "github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/jimlawless/whereami"

	"strconv"
	"strings"
	"time"
)

var monthMap = make(map[string]int)
var file *excelize.File

/**
 * Method: createSheet
 * Description: Create a new sheet with the passed in name using the medicine shoppe excel format.
 * Parameters:
 *    - acName The name of the new sheet
 */
func createSheet(acName string) int {
	// Create a new sheet.
	index := file.NewSheet(acName)

	style, err := file.NewStyle(`{"font":{"bold":true}}`)
	issue(err, whereami.WhereAmI())
	style2, err := file.NewStyle(`{"border":[{"type":"bottom","color":"#000000", "style":1}]}`)
	issue(err, whereami.WhereAmI())
	style3, err := file.NewStyle(`{"border":[{"type":"bottom","color":"#000000", "style":1},{"type":"top",
"color":"#000000", "style":1},{"type":"right","color":"#000000", "style":1},{"type":"left","color":"#000000", 
"style":1}]}`)
	issue(err, whereami.WhereAmI())

	// Set Header Styles
	issue(file.SetCellStyle(acName, "A1", "A4", style), whereami.WhereAmI())
	issue(file.SetCellStyle(acName, "F1", "F2", style), whereami.WhereAmI())
	issue(file.SetCellStyle(acName, "F5", "F5", style), whereami.WhereAmI())
	issue(file.SetCellStyle(acName, "B1", "B4", style2), whereami.WhereAmI())
	issue(file.SetCellStyle(acName, "C1", "C4", style2), whereami.WhereAmI())
	issue(file.SetCellStyle(acName, "D1", "D4", style2), whereami.WhereAmI())
	issue(file.SetCellStyle(acName, "G1", "I2", style2), whereami.WhereAmI())
	issue(file.SetCellStyle(acName, "G5", "G5", style3), whereami.WhereAmI())
	issue(file.SetCellStyle(acName, "G7", "J7", style3), whereami.WhereAmI())
	issue(file.SetCellStyle(acName, "A8", "J8", style3), whereami.WhereAmI())
	issue(file.SetCellStyle(acName, "G9", "J9", style3), whereami.WhereAmI())
	issue(file.SetCellStyle(acName, "I6", "J6", style3), whereami.WhereAmI())

	issue(file.MergeCell(acName, "A7", "C7"), whereami.WhereAmI())
	issue(file.MergeCell(acName, "D7", "F7"), whereami.WhereAmI())
	issue(file.MergeCell(acName, "A9", "F9"), whereami.WhereAmI())

	// Set value header cells
	issue(file.SetCellValue(acName, "A1", "DRUG NAME:"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "A2", "NDC:"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "A3", "FORM:"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "A4", "ITEM NUMBER:"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "F1", "DATE:"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "F2", "PKG SIZE:"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "F5", "STARTING:"), whereami.WhereAmI())

	issue(file.SetCellValue(acName, "A7", "PURCHASES"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "D7", "PRESCRIPTIONS"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "G7", "PROJECTED"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "H7", "ACTUAL"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "I6", "RPH"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "I7", "INITIALS"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "J6", "DATE"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "J7", "LOGGED"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "A8", "ORDER FORM #"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "B8", "DATE"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "C8", "QTY"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "D8", "PRESCRIPTION #"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "E8", "DATE"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "F8", "QTY"), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "A9", "STARTING INVENTORY"), whereami.WhereAmI())

	styleA, err := file.NewStyle(`{"alignment":{"horizontal":"center"}}`)
	issue(file.SetCellStyle(acName, "A7", "A7", styleA), whereami.WhereAmI())
	issue(file.SetCellStyle(acName, "D7", "D7", styleA), whereami.WhereAmI())
	issue(file.SetCellStyle(acName, "A9", "A9", styleA), whereami.WhereAmI())

	// Create the chart
	for i := 7; i < 3000; i++ {
		issue(file.SetCellStyle(acName, "A"+strconv.Itoa(i), "J"+strconv.Itoa(i), style3), whereami.WhereAmI())
		if i >= 10 {
			issue(file.SetCellFormula(acName, "G"+strconv.Itoa(i), "=G"+strconv.Itoa(i-1)+" + C"+strconv.Itoa(
				i)+"- F"+strconv.Itoa(i)), whereami.WhereAmI())
		}
	}

	return index
}

/**
 * Method: getSheet
 * Description: Get the newly made sheet, and populate it with data.
 * Parameters:
 *    - acNdc The ndc of the drug
 *    - acName The name of the sheet and drug
 */
func getSheet(acNdc string, acName string) {
	var lrQty float64
	var lcDate, lcLogDate time.Time
	var lcForm, lcItem, lcSize, lcPharm, lcScript, lcType string

	originalDate := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	programDate := time.Date(2019, 5, 3, 0, 0, 0, 0, time.UTC)

	issue(db.QueryRow("SELECT form, item_num, date, size from drugDB where ndc = $1", acNdc).Scan(&lcForm, &lcItem,
		&lcDate, &lcSize), whereami.WhereAmI())

	issue(file.SetCellValue(acName, "B1", acName), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "B2", acNdc), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "B3", lcForm), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "B4", lcItem), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "G1", strconv.Itoa(monthMap[lcDate.Month().String()])+"/"+strconv.Itoa(
		lcDate.Day())+"/"+strconv.Itoa(lcDate.Year())), whereami.WhereAmI())
	issue(file.SetCellValue(acName, "G2", lcSize), whereami.WhereAmI())

	rows, err := db.Query("SELECT pharmacist, qty, date, logdate, script, "+
		"type from orderDB where ndc = $1 order by date, id", acNdc)
	issue(err, whereami.WhereAmI())

	row := 10
	for rows.Next() {
		if rows.Err() != nil {
			issue(rows.Err(), whereami.WhereAmI())
			break
		}
		issue(rows.Scan(&lcPharm, &lrQty, &lcDate, &lcLogDate, &lcScript, &lcType), whereami.WhereAmI())

		if lcDate.Year() == originalDate.Year() && lcDate.Month() == originalDate.Month() && lcDate.Day() == originalDate.Day() {

			issue(file.SetCellValue(acName, "G5", lrQty), whereami.WhereAmI())
			issue(file.SetCellValue(acName, "G9", lrQty), whereami.WhereAmI())
			row--

		} else if strings.ToUpper(lcType) == "PURCHASE" {

			fillPurchase(acName, lcPharm, lcScript, lcDate, lrQty, lcLogDate, row)

		} else if strings.ToUpper(lcType) == "ACTUAL COUNT" {

			// TODO: Believe this is no longer needed but may need to reintroduce in testing
			if lcDate.Before(programDate) {

				issue(file.SetCellFormula(acName, "G"+strconv.Itoa(row), ""), whereami.WhereAmI())
				issue(file.SetCellValue(acName, "G"+strconv.Itoa(row), lrQty), whereami.WhereAmI())

			}

			fillCount(acName, lcPharm, lcType, lcDate, lrQty, lcLogDate, row)

		} else if strings.ToUpper(lcType) == "OVER/SHORT" {

			fillCount(acName, lcPharm, lcType, lcDate, lrQty, lcLogDate, row)

		} else if strings.ToUpper(lcType) == "REAL COUNT" {

			issue(file.SetCellFormula(acName, "G"+strconv.Itoa(row), ""), whereami.WhereAmI())
			issue(file.SetCellValue(acName, "G"+strconv.Itoa(row), lrQty), whereami.WhereAmI())
			fillCount(acName, lcPharm, lcType, lcDate, lrQty, lcLogDate, row)

		} else if strings.ToUpper(lcType) == "AUDIT" {

			issue(file.SetCellFormula(acName, "G"+strconv.Itoa(row), "=G"+strconv.Itoa(row-2)+" + C"+strconv.Itoa(
				row-1)+"- F"+strconv.Itoa(row-1)), whereami.WhereAmI())
			issue(file.SetCellValue(acName, "G"+strconv.Itoa(row), lrQty), whereami.WhereAmI())
			fillCount(acName, lcPharm, lcType, lcDate, lrQty, lcLogDate, row)

		} else {

			fillScript(acName, lcPharm, lcScript, lcDate, lrQty, lcLogDate, row)

		}

		row++

	}

	issue(rows.Close(), whereami.WhereAmI())
}

/**
 * Method: fillPurchase
 * Description: Fill a purchase line using the passed in pieces.
 * Parameters: All the inputs to print on line 'row'.
 */
func fillPurchase(acName string, acPharm string, acScript string, acDate time.Time, arQty float64,
	acLogDate time.Time, row int) {
	issue(file.SetCellValue(acName, "A"+strconv.Itoa(row), acScript), whereami.WhereAmI())

	issue(file.SetCellValue(acName, "B"+strconv.Itoa(row), strconv.Itoa(monthMap[acDate.Month().String()])+"/"+strconv.Itoa(acDate.Day())+"/"+strconv.Itoa(acDate.Year())), whereami.WhereAmI())

	issue(file.SetCellValue(acName, "C"+strconv.Itoa(row), arQty), whereami.WhereAmI())

	issue(file.SetCellValue(acName, "I"+strconv.Itoa(row), acPharm), whereami.WhereAmI())

	issue(file.SetCellValue(acName, "J"+strconv.Itoa(row), strconv.Itoa(monthMap[acLogDate.Month().String()])+"/"+strconv.Itoa(acLogDate.Day())+"/"+strconv.Itoa(acLogDate.Year())), whereami.WhereAmI())
}

/**
 * Method: fillScript
 * Description: Fill a script line using the passed in pieces.
 * Parameters: All the inputs to print on line 'row'.
 */
func fillScript(acName string, acPharm string, acScript string, acDate time.Time, arQty float64,
	acLogDate time.Time, row int) {
	issue(file.SetCellValue(acName, "D"+strconv.Itoa(row), acScript), whereami.WhereAmI())

	issue(file.SetCellValue(acName, "E"+strconv.Itoa(row), strconv.Itoa(monthMap[acDate.Month().String()])+"/"+strconv.Itoa(acDate.Day())+"/"+strconv.Itoa(acDate.Year())), whereami.WhereAmI())

	issue(file.SetCellValue(acName, "F"+strconv.Itoa(row), arQty), whereami.WhereAmI())

	issue(file.SetCellValue(acName, "I"+strconv.Itoa(row), acPharm), whereami.WhereAmI())

	issue(file.SetCellValue(acName, "J"+strconv.Itoa(row), strconv.Itoa(monthMap[acLogDate.Month().String()])+"/"+strconv.Itoa(acLogDate.Day())+"/"+strconv.Itoa(acLogDate.Year())), whereami.WhereAmI())
}

/**
 * Method: fillCount
 * Description: Fill a real or actual count line using the passed in pieces.
 * Parameters: All the inputs to print on line 'row'.
 */
func fillCount(acName string, acPharm string, acType string, acDate time.Time, arQty float64,
	acLogDate time.Time, row int) {

	style, err := file.NewStyle(`{"fill":{"type":"gradient","color":["#FFFF00","#FFFF00"],"shading":1}}`)
	issue(err, whereami.WhereAmI())

	issue(file.SetCellStyle(acName, "F"+strconv.Itoa(row), "F"+strconv.Itoa(row), style), whereami.WhereAmI())

	// The use the same columns
	fillScript(acName, acPharm, acType, acDate, arQty, acLogDate, row)
}

/**
 * Method: fixName
 * Description: Ensure that the name is properly formatted for excel.
 * Parameters:
 *    -acName The name to fix
 *    -acSeenMap The map of seen names and counts
 */
func fixName(acName string, acSeenMap map[string]int) string {

	if strings.Contains(acName, "/") {
		acName = strings.ReplaceAll(acName, "/", "-")
	}

	if strings.Contains(acName, "(") && strings.Contains(acName, ")") {
		acName = string([]rune(acName)[0 : len(acName)-3])
	}

	if len(acName) > 30 {
		acName = strings.TrimSpace(string([]rune(acName)[0:30]))
	}

	lcName := strings.ToUpper(acName)

	if _, ok := acSeenMap[lcName]; ok {
		acName = acName + " (" + strconv.Itoa(acSeenMap[lcName]) + ")"
		acSeenMap[lcName]++
	} else {
		acSeenMap[lcName] = 1
	}

	return acName
}

/**
 * Method: ExcelWriter
 * Description: Write the database to an excel document.
 * Parameter:
 *    - acFileName The name to save the file to.
 */
func ExcelWriter(acFileName string) {
	var lcName, lcNdc string

	file = excelize.NewFile()

	Log("Writing database to excel format in file "+acFileName, INFO, whereami.WhereAmI())

	seenMap := make(map[string]int)

	monthMap["January"] = 1
	monthMap["February"] = 2
	monthMap["March"] = 3
	monthMap["April"] = 4
	monthMap["May"] = 5
	monthMap["June"] = 6
	monthMap["July"] = 7
	monthMap["August"] = 8
	monthMap["September"] = 9
	monthMap["October"] = 10
	monthMap["November"] = 11
	monthMap["December"] = 12

	rows, err := db.Query("SELECT name, ndc from drugDB order by name")
	issue(err, whereami.WhereAmI())

	var lnCounter int
	lnCounter = 0
	for rows.Next() {
		if rows.Err() != nil {
			issue(rows.Err(), whereami.WhereAmI())
			break
		}

		issue(rows.Scan(&lcName, &lcNdc), whereami.WhereAmI())
		lcName = fixName(lcName, seenMap)
		createSheet(lcName)
		getSheet(lcNdc, lcName)
		lnCounter++
	}

	Log("Wrote "+strconv.Itoa(lnCounter)+" drugs to excel", DEBUG, whereami.WhereAmI())

	// Set active sheet of the workbook.
	file.SetActiveSheet(0)

	file.DeleteSheet("Sheet1")

	// Save xlsx file by the given path.
	issue(file.SaveAs(acFileName+".xlsx"), whereami.WhereAmI())
	Log("Excel file closed", DEBUG, whereami.WhereAmI())

	issue(rows.Close(), whereami.WhereAmI())
}
