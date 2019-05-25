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
@param ndc The NDC code
@param str The current error string
@return The current error string
*/
func CheckNDC(ndc string, str string) (string, string) {
	if len(ndc) != 11 && len(ndc) != 13 {
		str = "NDC is not the correct length"
	} else if ndc[5] != '-' || ndc[10] != '-' {
		ndc = ndc[:5] + "-" + ndc[5:9] + "-" + ndc[9:]
	}
	return ndc, str
}

/**
Function: CheckDate
Description: Checks to make sure all parts of the date are the
correct length and ints
@param month The month
@param day The day
@param year The year
@param str The current error string
@return The current error string
*/
func CheckDate(month string, day string, year string, str string) string {
	if len(month) != 2 && len(month) != 1 {
		str = "Month must be in the format XX or X"
	} else if len(day) != 2 && len(day) != 1 {
		str = "Day must be in the format XX or X"
	} else if len(year) != 4 {
		str = "Year must be in the format XXXX"
	}
	_, err := strconv.Atoi(month)
	if err != nil {
		str = "The month entered was not a number"
	}
	_, err = strconv.Atoi(day)
	if err != nil {
		str = "The day entered was not a number"
	}
	_, err = strconv.Atoi(year)
	if err != nil {
		str = "The year entered was not a number"
	}

	return str
}

/**
Function: CheckQty
Description: Checks if the quantity is greater than 0
@param qty The quantity entered
@param str The current error string
@return The current error string
*/
func CheckQty(qty string, str string) string {
	qt, _ := strconv.Atoi(qty)
	if qt < 0 {
		str = "Quantity must be greater than 0"
	}
	return str
}

/**
Function: CheckNum
Description: Checks if the entered value was a number
@param number The supposed number
@param str The current error string
@return The current error string
*/
//func CheckNum(number string, str string) string {
//	_, err := strconv.Atoi(number)
//	if err != nil {
//		if number == "" {
//			str = "You missed a number field!"
//		} else {
//			str = number + " is not a valid number"
//		}
//	}
//	return str
//}

/**
Function: CheckString
Description: Checks if the entered value was a non empty string
@param input The supposed string
@param str The current error string
@return The current error string
*/
func CheckString(input string, str string) string {
	if input == "" {
		str = "You missed a text field!"
	}
	return str
}
