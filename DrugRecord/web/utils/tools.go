package utils

import (
	"strings"
)

/**
Function: ParseDate
Description: Takes in a date in the form MM DD YYYY as a single string and returns it parsed.
This method excepts dates broken up by - . or /
@param acDate The date string to parse
@return The date in the order month, day, year as a string
*/
func ParseDate(acDate string) (string, string, string) {
	if strings.Contains(acDate, "/") {
		lcSplitString := strings.Split(acDate, "/")
		return lcSplitString[0], lcSplitString[1], lcSplitString[2]
	} else if strings.Contains(acDate, ".") {
		lcSplitString := strings.Split(acDate, ".")
		return lcSplitString[0], lcSplitString[1], lcSplitString[2]
	} else if strings.Contains(acDate, "-") {
		lcSplitString := strings.Split(acDate, "-")
		return lcSplitString[0], lcSplitString[1], lcSplitString[2]
	} else {
		return "", "", ""
	}
}
