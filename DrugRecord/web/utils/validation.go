/**
File: validation.go
Description: Checks a multitude of different input for validity.
@author Bryan Conn
@date 11/26/2018
*/
package utils

import "strconv"

/**
Function: checkNDC
Description: Checks to make sure that the NDC is the correct
length and has dashes in the right spots.
@param acNdc The NDC code
@param acErrorString The current error string
@return The current error string
*/
func CheckNDC(acNdc string, acErrorString string) (string, string) {
	if len(acNdc) != 11 && len(acNdc) != 13 {
		acErrorString = "NDC is not the correct length"
	} else if acNdc[5] != '-' || acNdc[10] != '-' {
		acNdc = acNdc[:5] + "-" + acNdc[5:9] + "-" + acNdc[9:]
	}
	return acNdc, acErrorString
}

/**
Function: CheckDate
Description: Checks to make sure all parts of the date are the
correct length and ints
@param acMonth The month
@param acDay The day
@param acYear The year
@param acErrorString The current error string
@return The current error string
*/
func CheckDate(acMonth string, acDay string, acYear string, acErrorString string) (string, string) {
	if len(acMonth) != 2 && len(acMonth) != 1 {
		acErrorString = "Month must be in the format XX or X"
	} else if len(acDay) != 2 && len(acDay) != 1 {
		acErrorString = "Day must be in the format XX or X"
	} else if len(acYear) != 4 && len(acYear) != 2 {
		acErrorString = "Year must be in the format XXXX or XX"
	}
	_, err := strconv.Atoi(acMonth)
	if err != nil {
		acErrorString = "The month entered was not a number"
	}
	_, err = strconv.Atoi(acDay)
	if err != nil {
		acErrorString = "The day entered was not a number"
	}

	if len(acYear) == 2 {
		acYear = "20" + acYear
	}

	_, err = strconv.Atoi(acYear)
	if err != nil {
		acErrorString = "The year entered was not a number"
	}

	return acErrorString, acYear
}

/**
Function: CheckQty
Description: Checks if the quantity is greater than 0
@param qty The quantity entered
@param acErrorString The current error string
@return The current error string
*/
func CheckQty(acQty string, acErrorString string) string {
	lnQty, _ := strconv.Atoi(acQty)
	if lnQty < 0 {
		acErrorString = "Quantity must be greater than 0"
	}
	return acErrorString
}

/**
Function: CheckNum
Description: Checks if the entered value was a number
@param number The supposed number
@param str The current error string
@return The current error string
*/
// func CheckNum(number string, str string) string {
//	_, err := strconv.Atoi(number)
//	if err != nil {
//		if number == "" {
//			str = "You missed a number field!"
//		} else {
//			str = number + " is not a valid number"
//		}
//	}
//	return str
// }

/**
Function: CheckString
Description: Checks if the entered value was a non empty string
@param acInput The supposed string
@param str The current error string
@return The current error string
*/
func CheckString(acInput string, acErrorString string) string {
	if acInput == "" {
		acErrorString = "You missed a text field!"
	}
	return acErrorString
}
