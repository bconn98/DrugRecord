package webUtils

import (
	"strconv"
	"strings"

	"github.com/jimlawless/whereami"

	"github.com/bconn98/DrugRecord/utils"
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

func ParseFloat(acFloatVal string) float64 {
	lrQty, err := strconv.ParseFloat(acFloatVal, 64)
	if err != nil {
		utils.Log(err.Error(), utils.ERROR, whereami.WhereAmI())
	}
	return lrQty
}
