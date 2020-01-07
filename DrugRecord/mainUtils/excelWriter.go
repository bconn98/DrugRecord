package mainUtils

import (
	"github.com/360EntSecGroup-Skylar/excelize"
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
	issue(err)
	style2, err := file.NewStyle(`{"border":[{"type":"bottom","color":"#000000", "style":1}]}`)
	issue(err)
	style3, err := file.NewStyle(`{"border":[{"type":"bottom","color":"#000000", "style":1},{"type":"top",
"color":"#000000", "style":1},{"type":"right","color":"#000000", "style":1},{"type":"left","color":"#000000", 
"style":1}]}`)
	issue(err)

	// Set Header Styles
	issue(file.SetCellStyle(acName, "A1", "A4", style))
	issue(file.SetCellStyle(acName, "F1", "F2", style))
	issue(file.SetCellStyle(acName, "F5", "F5", style))
	issue(file.SetCellStyle(acName, "B1", "B4", style2))
	issue(file.SetCellStyle(acName, "C1", "C4", style2))
	issue(file.SetCellStyle(acName, "D1", "D4", style2))
	issue(file.SetCellStyle(acName, "G1", "I2", style2))
	issue(file.SetCellStyle(acName, "G5", "G5", style3))
	issue(file.SetCellStyle(acName, "G7", "J7", style3))
	issue(file.SetCellStyle(acName, "A8", "J8", style3))
	issue(file.SetCellStyle(acName, "G9", "J9", style3))
	issue(file.SetCellStyle(acName, "I6", "J6", style3))

	issue(file.MergeCell(acName, "A7", "C7"))
	issue(file.MergeCell(acName, "D7", "F7"))
	issue(file.MergeCell(acName, "A9", "F9"))

	// Set value header cells
	issue(file.SetCellValue(acName, "A1", "DRUG NAME:"))
	issue(file.SetCellValue(acName, "A2", "NDC:"))
	issue(file.SetCellValue(acName, "A3", "FORM:"))
	issue(file.SetCellValue(acName, "A4", "ITEM NUMBER:"))
	issue(file.SetCellValue(acName, "F1", "DATE:"))
	issue(file.SetCellValue(acName, "F2", "PKG SIZE:"))
	issue(file.SetCellValue(acName, "F5", "STARTING:"))

	issue(file.SetCellValue(acName, "A7", "PURCHASES"))
	issue(file.SetCellValue(acName, "D7", "PRESCRIPTIONS"))
	issue(file.SetCellValue(acName, "G7", "PROJECTED"))
	issue(file.SetCellValue(acName, "H7", "ACTUAL"))
	issue(file.SetCellValue(acName, "I6", "RPH"))
	issue(file.SetCellValue(acName, "I7", "INITIALS"))
	issue(file.SetCellValue(acName, "J6", "DATE"))
	issue(file.SetCellValue(acName, "J7", "LOGGED"))
	issue(file.SetCellValue(acName, "A8", "ORDER FORM #"))
	issue(file.SetCellValue(acName, "B8", "DATE"))
	issue(file.SetCellValue(acName, "C8", "QTY"))
	issue(file.SetCellValue(acName, "D8", "PRESCRIPTION #"))
	issue(file.SetCellValue(acName, "E8", "DATE"))
	issue(file.SetCellValue(acName, "F8", "QTY"))
	issue(file.SetCellValue(acName, "A9", "STARTING INVENTORY"))

	styleA, err := file.NewStyle(`{"alignment":{"horizontal":"center"}}`)
	issue(file.SetCellStyle(acName, "A7", "A7", styleA))
	issue(file.SetCellStyle(acName, "D7", "D7", styleA))
	issue(file.SetCellStyle(acName, "A9", "A9", styleA))

	// Create the chart
	for i := 7; i < 3000; i++ {
		issue(file.SetCellStyle(acName, "A"+strconv.Itoa(i), "J"+strconv.Itoa(i), style3))
		if i >= 10 {
			issue(file.SetCellFormula(acName, "G"+strconv.Itoa(i), "=G"+strconv.Itoa(i-1)+" + C"+strconv.Itoa(
				i)+"- F"+strconv.Itoa(i)))
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
		&lcDate, &lcSize))

	issue(file.SetCellValue(acName, "B1", acName))
	issue(file.SetCellValue(acName, "B2", acNdc))
	issue(file.SetCellValue(acName, "B3", lcForm))
	issue(file.SetCellValue(acName, "B4", lcItem))
	issue(file.SetCellValue(acName, "G1", strconv.Itoa(monthMap[lcDate.Month().String()])+"/"+strconv.Itoa(
		lcDate.Day())+"/"+strconv.Itoa(lcDate.Year())))
	issue(file.SetCellValue(acName, "G2", lcSize))

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

			issue(file.SetCellValue(acName, "G5", lrQty))
			issue(file.SetCellValue(acName, "G9", lrQty))
			row--

		} else if strings.ToUpper(lcType) == "PURCHASE" {

			fillPurchase(acName, lcPharm, lcScript, lcDate, lrQty, lcLogDate, row)

		} else if strings.ToUpper(lcType) == "ACTUAL COUNT" {

			// TODO: Believe this is no longer needed but may need to reintroduce in testing
			if lcDate.Before(programDate) {

				issue(file.SetCellFormula(acName, "G"+strconv.Itoa(row), ""))
				issue(file.SetCellValue(acName, "G"+strconv.Itoa(row), lrQty))

			}

			fillCount(acName, lcPharm, lcType, lcDate, lrQty, lcLogDate, row)

		} else if strings.ToUpper(lcType) == "OVER/SHORT" {

			fillCount(acName, lcPharm, lcType, lcDate, lrQty, lcLogDate, row)

		} else if strings.ToUpper(lcType) == "REAL COUNT" {

			issue(file.SetCellFormula(acName, "G"+strconv.Itoa(row), ""))
			issue(file.SetCellValue(acName, "G"+strconv.Itoa(row), lrQty))
			fillCount(acName, lcPharm, lcType, lcDate, lrQty, lcLogDate, row)

		} else if strings.ToUpper(lcType) == "AUDIT" {

			issue(file.SetCellFormula(acName, "G"+strconv.Itoa(row), "=G"+strconv.Itoa(row-2)+" + C"+strconv.Itoa(
				row-1)+"- F"+strconv.Itoa(row-1)))
			issue(file.SetCellValue(acName, "G"+strconv.Itoa(row), lrQty))
			fillCount(acName, lcPharm, lcType, lcDate, lrQty, lcLogDate, row)

		} else {

			fillScript(acName, lcPharm, lcScript, lcDate, lrQty, lcLogDate, row)

		}

		row++

	}

	issue(rows.Close())
}

/**
 * Method: fillPurchase
 * Description: Fill a purchase line using the passed in pieces.
 * Parameters: All the inputs to print on line 'row'.
 */
func fillPurchase(acName string, acPharm string, acScript string, acDate time.Time, arQty float64,
	acLogDate time.Time, row int) {
	issue(file.SetCellValue(acName, "A"+strconv.Itoa(row), acScript))

	issue(file.SetCellValue(acName, "B"+strconv.Itoa(row), strconv.Itoa(monthMap[acDate.Month().String()])+"/"+strconv.Itoa(acDate.Day())+"/"+strconv.Itoa(acDate.Year())))

	issue(file.SetCellValue(acName, "C"+strconv.Itoa(row), arQty))

	issue(file.SetCellValue(acName, "I"+strconv.Itoa(row), acPharm))

	issue(file.SetCellValue(acName, "J"+strconv.Itoa(row), strconv.Itoa(monthMap[acLogDate.Month().String()])+"/"+strconv.Itoa(acLogDate.Day())+"/"+strconv.Itoa(acLogDate.Year())))
}

/**
 * Method: fillScript
 * Description: Fill a script line using the passed in pieces.
 * Parameters: All the inputs to print on line 'row'.
 */
func fillScript(acName string, acPharm string, acScript string, acDate time.Time, arQty float64,
	acLogDate time.Time, row int) {
	issue(file.SetCellValue(acName, "D"+strconv.Itoa(row), acScript))

	issue(file.SetCellValue(acName, "E"+strconv.Itoa(row), strconv.Itoa(monthMap[acDate.Month().String()])+"/"+strconv.Itoa(acDate.Day())+"/"+strconv.Itoa(acDate.Year())))

	issue(file.SetCellValue(acName, "F"+strconv.Itoa(row), arQty))

	issue(file.SetCellValue(acName, "I"+strconv.Itoa(row), acPharm))

	issue(file.SetCellValue(acName, "J"+strconv.Itoa(row), strconv.Itoa(monthMap[acLogDate.Month().String()])+"/"+strconv.Itoa(acLogDate.Day())+"/"+strconv.Itoa(acLogDate.Year())))
}

/**
 * Method: fillCount
 * Description: Fill a real or actual count line using the passed in pieces.
 * Parameters: All the inputs to print on line 'row'.
 */
func fillCount(acName string, acPharm string, acType string, acDate time.Time, arQty float64,
	acLogDate time.Time, row int) {

	style, err := file.NewStyle(`{"fill":{"type":"gradient","color":["#FFFF00","#FFFF00"],"shading":1}}`)
	issue(err)

	issue(file.SetCellStyle(acName, "F"+strconv.Itoa(row), "F"+strconv.Itoa(row), style))

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

	for rows.Next() {
		if rows.Err() != nil {
			issue(rows.Err())
			break
		}

		issue(rows.Scan(&lcName, &lcNdc))
		lcName = fixName(lcName, seenMap)
		createSheet(lcName)
		getSheet(lcNdc, lcName)
	}

	// Set active sheet of the workbook.
	file.SetActiveSheet(0)

	file.DeleteSheet("Sheet1")

	// Save xlsx file by the given path.
	issue(file.SaveAs(acFileName + ".xlsx"))

	issue(rows.Close())
}
