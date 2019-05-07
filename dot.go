package humanize

import (
	"fmt"
	"strings"
	"strconv"
)

// Dot gets the string with comma and transforms in to a string with dots
// just that simple so us from Brazil don't feel excluded
// e.g 834,142 -> 834.142
func Dot(v int64) string {
	return strings.Replace(Comma(v), ",", ".", -1)
}

// Dotf gets the string with comma and transforms in to a string with dots
// just that simple so us from Brazil don't feel excluded
// e.g 834,142.30 -> 834.142,30
func Dotf(v float64) string {
	// spliting the number in number after and before dot: "100000.00" -> "100000" & ".00"
	var vSplitInDot = strings.Split(fmt.Sprintf("%f", v), ".")

	// get the value befor the dot with commmas
	var vBeforeDot, _ = strconv.ParseFloat(vSplitInDot[0], 64)
	var vWithCommas = Commaf(vBeforeDot)

	// and here replace the commas with dots of the Commaf result: 100,000 -> 100.000
	var vCommasToDots = strings.Replace(vWithCommas, ",", ".", -1)
	// return the values with dots finnaly: "100.000" & ".00" -> "100.000,00"
	return fmt.Sprintf("%s,%s", vCommasToDots, vSplitInDot[1][0:])
}

// DotWithDigits get the float and put the number of
// decimals as in the argument d
// e.g  v = 834,142.30 and d = 1 -> 834,142.3
func DotWithDigits(v float64, d int) string {
	return stripTrailingDigits(Dotf(v), d)
}
