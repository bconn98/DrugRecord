package mainUtils

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
	"strings"
	"time"
)

var monthMap = make(map[string]int)
var file *excelize.File
var blankStyle int
var borderStyle int
var boldStyle int
var underlineStyle int
var highlightStyle int
var blankCell excelize.Cell
var borderCell excelize.Cell

/**
 * Method: createSheet
 * Description: Create a new sheet with the passed in name using the medicine shoppe excel format.
 * Parameters:
 *    - acName The name of the new sheet
 */
func createSheet(acName string) int {
	// Create a new sheet.
	index := file.NewSheet(acName)

	issue(file.MergeCell(acName, "A7", "C7"))
	issue(file.MergeCell(acName, "D7", "F7"))
	issue(file.MergeCell(acName, "A9", "F9"))

	// Create the chart
	for i := 7; i < 3000; i++ {
		if i >= 10 {
			issue(file.SetCellFormula(acName, "G"+strconv.Itoa(i), "=G"+strconv.Itoa(i-1)+" + C"+strconv.Itoa(
				i)+"- F"+strconv.Itoa(i)))
		}
	}

	return index
}

func makeHeader(acStreamWriter *excelize.StreamWriter, acName string, acNdc string, acForm string, acItem string,
	acDate string, acSize string) {

	var underlineCell = excelize.Cell{StyleID: underlineStyle, Value: ""}

	issue(acStreamWriter.SetRow("A1", []interface{}{ // Row 1
		excelize.Cell{StyleID: boldStyle, Value: "DRUG NAME:"}, // A
		excelize.Cell{StyleID: underlineStyle, Value: acName},  // B
		underlineCell, // C
		underlineCell, // D
		blankCell,     // E
		excelize.Cell{StyleID: boldStyle, Value: "DATE:"},     // F
		excelize.Cell{StyleID: underlineStyle, Value: acDate}, // G
		underlineCell,   // H
		underlineCell})) // I

	issue(acStreamWriter.SetRow("A2", []interface{}{ // Row 2
		excelize.Cell{StyleID: boldStyle, Value: "NDC:"},     // A
		excelize.Cell{StyleID: underlineStyle, Value: acNdc}, // B
		underlineCell, // C
		underlineCell, // D
		blankCell,     // E
		excelize.Cell{StyleID: boldStyle, Value: "PKG SIZE:"}, // F
		excelize.Cell{StyleID: underlineStyle, Value: acSize}, // G
		underlineCell,   // H
		underlineCell})) // I

	issue(acStreamWriter.SetRow("A3", []interface{}{ // Row 3
		excelize.Cell{StyleID: boldStyle, Value: "FORM:"},     // A
		excelize.Cell{StyleID: underlineStyle, Value: acForm}, // B
		underlineCell,   // C
		underlineCell})) // D

	issue(acStreamWriter.SetRow("A4", []interface{}{ // Row 4
		excelize.Cell{StyleID: boldStyle, Value: "ITEM NUMBER:"}, // A
		excelize.Cell{StyleID: underlineStyle, Value: acItem},    // B
		underlineCell,   // C
		underlineCell})) // D

	issue(acStreamWriter.SetRow("A6", []interface{}{ // Row 6
		blankCell, // A
		blankCell, // B
		blankCell, // C
		blankCell, // D
		blankCell, // E
		blankCell, // F
		blankCell, // G
		blankCell, // H
		excelize.Cell{StyleID: borderStyle, Value: "RPH"},    // I
		excelize.Cell{StyleID: borderStyle, Value: "DATE"}})) // J

	issue(acStreamWriter.SetRow("A7", []interface{}{ // Row 7
		excelize.Cell{StyleID: borderStyle, Value: "PURCHASES"}, // A
		borderCell, // B
		borderCell, // C
		excelize.Cell{StyleID: borderStyle, Value: "PRESCRIPTIONS"}, // D
		borderCell, // E
		borderCell, // F
		excelize.Cell{StyleID: borderStyle, Value: "PROJECTED"}, // G
		excelize.Cell{StyleID: borderStyle, Value: "ACTUAL"},    // H
		excelize.Cell{StyleID: borderStyle, Value: "INITIALS"},  // I
		excelize.Cell{StyleID: borderStyle, Value: "LOGGED"},    // J
	}))

	issue(acStreamWriter.SetRow("A8", []interface{}{ // Row 8
		excelize.Cell{StyleID: borderStyle, Value: "ORDER FORM #"},   // A
		excelize.Cell{StyleID: borderStyle, Value: "DATE"},           // B
		excelize.Cell{StyleID: borderStyle, Value: "QTY"},            // C
		excelize.Cell{StyleID: borderStyle, Value: "PRESCRIPTION #"}, // D
		excelize.Cell{StyleID: borderStyle, Value: "DATE"},           // E
		excelize.Cell{StyleID: borderStyle, Value: "QTY"},            // F
		borderCell,   // G
		borderCell,   // H
		borderCell,   // I
		borderCell})) // J
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
		&lcDate, &lcSize))

	streamWriter, err := file.NewStreamWriter(acName)
	issue(err)

	date := strconv.Itoa(monthMap[lcDate.Month().String()]) + "/" + strconv.Itoa(
		lcDate.Day()) + "/" + strconv.Itoa(lcDate.Year())

	makeHeader(streamWriter, acName, acNdc, lcForm, lcItem, date, lcSize)

	rows, err := db.Query("SELECT pharmacist, qty, date, logdate, script, "+
		"type from orderDB where ndc = $1 order by date, id", acNdc)
	issue(err)

	row := 10
	for rows.Next() {
		if rows.Err() != nil {
			issue(rows.Err())
			break
		}
		issue(rows.Scan(&lcPharm, &lrQty, &lcDate, &lcLogDate, &lcScript, &lcType))

		if lcDate.Year() == originalDate.Year() && lcDate.Month() == originalDate.Month() && lcDate.Day() == originalDate.Day() {

			issue(streamWriter.SetRow("A5", []interface{}{ // Row 5
				blankCell, // A
				blankCell, // B
				blankCell, // C
				blankCell, // D
				blankCell, // E
				excelize.Cell{StyleID: boldStyle, Value: "STARTING:"}, // F
				excelize.Cell{StyleID: borderStyle, Value: lrQty}}))   // G

			issue(streamWriter.SetRow("A9", []interface{}{ // Row 9
				excelize.Cell{StyleID: borderStyle, Value: "STARTING INVENTORY"}, // A
				blankCell, // B
				blankCell, // C
				blankCell, // D
				blankCell, // E
				blankCell, // F
				excelize.Cell{StyleID: borderStyle, Value: lrQty}, // G
				borderCell,   // H
				borderCell,   // I
				borderCell})) // J
			row-- // Negate row++ below

		} else if strings.ToUpper(lcType) == "PURCHASE" {

			fillPurchase(streamWriter, lcPharm, lcScript, lcDate, lrQty, lcLogDate, row)

		} else if strings.ToUpper(lcType) == "ACTUAL COUNT" {

			// TODO: Believe this is no longer needed but may need to reintroduce in testing
			if lcDate.Before(programDate) {

				// No formula for projected, hard code the value (lrQty)

			}

			fillCount(streamWriter, acName, lcPharm, lcType, lcDate, lrQty, lcLogDate, row)

		} else if strings.ToUpper(lcType) == "OVER/SHORT" { // TODO: Believe this was removed

			fillCount(streamWriter, acName, lcPharm, lcType, lcDate, lrQty, lcLogDate, row)

		} else if strings.ToUpper(lcType) == "REAL COUNT" {

			// No formula for projected, hard code the value (lrQty)

			fillCount(streamWriter, acName, lcPharm, lcType, lcDate, lrQty, lcLogDate, row)

		} else if strings.ToUpper(lcType) == "AUDIT" {

			// Use row above formula (G (row - 2) + C (row - 1) - F (row - 1)) projected,
			// hard code the value for row as (lrQty)

			fillCount(streamWriter, acName, lcPharm, lcType, lcDate, lrQty, lcLogDate, row)

		} else {

			fillScript(streamWriter, lcPharm, lcScript, lcDate, lrQty, lcLogDate, row)

		}

		row++

	}

	issue(streamWriter.Flush())
	issue(rows.Close())
}

/**
 * Method: fillPurchase
 * Description: Fill a purchase line using the passed in pieces.
 * Parameters: All the inputs to print on line 'row'.
 */
func fillPurchase(acStreamWriter *excelize.StreamWriter, acPharm string, acScript string, acDate time.Time,
	arQty float64,
	acLogDate time.Time, row int) {

	issue(acStreamWriter.SetRow("A"+strconv.Itoa(row), []interface{}{
		excelize.Cell{StyleID: borderStyle, Value: acScript}, // A
		excelize.Cell{StyleID: borderStyle, Value: strconv.Itoa(monthMap[acDate.Month().String()]) + "/" + strconv.Itoa(
			acDate.Day()) + "/" + strconv.Itoa(acDate.Year())}, // B
		excelize.Cell{StyleID: borderStyle, Value: arQty}, // C
		borderCell, // D
		borderCell, // E
		borderCell, // F
		borderCell, // G
		borderCell, // H
		excelize.Cell{StyleID: borderStyle, Value: acPharm}, // I
		excelize.Cell{StyleID: borderStyle, Value: strconv.Itoa(monthMap[acLogDate.Month().String()]) + "/" + strconv.Itoa(
			acLogDate.Day()) + "/" + strconv.Itoa(acLogDate.Year())}})) // J
}

/**
 * Method: fillScript
 * Description: Fill a script line using the passed in pieces.
 * Parameters: All the inputs to print on line 'row'.
 */
func fillScript(acStreamWriter *excelize.StreamWriter, acPharm string, acScript string, acDate time.Time, arQty float64,
	acLogDate time.Time, row int) {

	issue(acStreamWriter.SetRow("A"+strconv.Itoa(row), []interface{}{
		borderCell, // A
		borderCell, // B
		borderCell, // C
		excelize.Cell{StyleID: borderStyle, Value: acScript}, // D
		excelize.Cell{StyleID: borderStyle, Value: strconv.Itoa(monthMap[acDate.Month().String()]) + "/" + strconv.Itoa(
			acDate.Day()) + "/" + strconv.Itoa(acDate.Year())}, // E
		excelize.Cell{StyleID: borderStyle, Value: arQty}, // F
		borderCell, // G
		borderCell, // H
		excelize.Cell{StyleID: borderStyle, Value: acPharm}, // I
		excelize.Cell{StyleID: borderStyle, Value: strconv.Itoa(monthMap[acLogDate.Month().String()]) + "/" + strconv.Itoa(
			acLogDate.Day()) + "/" + strconv.Itoa(acLogDate.Year())}})) // J
}

/**
 * Method: fillCount
 * Description: Fill a real or actual count line using the passed in pieces.
 * Parameters: All the inputs to print on line 'row'.
 */
func fillCount(acStreamWriter *excelize.StreamWriter, acName string, acPharm string, acType string, acDate time.Time,
	arQty float64,
	acLogDate time.Time, row int) {

	issue(acStreamWriter.SetRow("A"+strconv.Itoa(row), []interface{}{
		borderCell, // A
		borderCell, // B
		borderCell, // C
		excelize.Cell{StyleID: borderStyle, Value: acType}, // D
		excelize.Cell{StyleID: borderStyle, Value: strconv.Itoa(monthMap[acDate.Month().String()]) + "/" + strconv.Itoa(
			acDate.Day()) + "/" + strconv.Itoa(acDate.Year())}, // E
		excelize.Cell{StyleID: highlightStyle, Value: arQty}, // F
		borderCell, // G
		borderCell, // H
		excelize.Cell{StyleID: borderStyle, Value: acPharm}, // I
		excelize.Cell{StyleID: borderStyle, Value: strconv.Itoa(monthMap[acLogDate.Month().String()]) + "/" + strconv.Itoa(
			acLogDate.Day()) + "/" + strconv.Itoa(acLogDate.Year())}})) // J
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
	issue(err)

	blankStyle, err = file.NewStyle(`{"font":{"color":"#000000"}}`)
	issue(err)
	borderStyle, err = file.NewStyle(`{"border":[{"type":"bottom","color":"#000000", "style":1},{"type":"top",
"color":"#000000", "style":1},{"type":"right","color":"#000000", "style":1},{"type":"left","color":"#000000", 
"style":1}]}`)
	issue(err)
	boldStyle, err = file.NewStyle(`{"font":{"bold":true}}`)
	underlineStyle, err = file.NewStyle(`{"border":[{"type":"bottom","color":"#000000", "style":1}]}`)
	issue(err)
	highlightStyle, err = file.NewStyle(`{"fill":{"type":"gradient","color":["#FFFF00","#FFFF00"],"shading":1},
"border":[{"type":"bottom","color":"#000000", "style":1},{"type":"top",
"color":"#000000", "style":1},{"type":"right","color":"#000000", "style":1},{"type":"left","color":"#000000", 
"style":1}]}`)
	issue(err)

	blankCell = excelize.Cell{StyleID: blankStyle, Value: ""}
	borderCell = excelize.Cell{StyleID: borderStyle, Value: ""}

	for rows.Next() {
		if rows.Err() != nil {
			issue(rows.Err())
			break
		}

		issue(rows.Scan(&lcName, &lcNdc))
		lcName = fixName(lcName, seenMap)
		file.NewSheet(lcName)

		getSheet(lcNdc, lcName)
	}

	// Set active sheet of the workbook.
	file.SetActiveSheet(0)

	file.DeleteSheet("Sheet1")

	// Save xlsx file by the given path.
	if strings.Contains(acFileName, ".xlsx") {
		issue(file.SaveAs(acFileName))
	} else {
		issue(file.SaveAs(acFileName + ".xlsx"))
	}

	issue(rows.Close())
}
