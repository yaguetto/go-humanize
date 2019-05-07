package humanize

import (
	"fmt"
	"strings"
)

// Dot gets the string with comma and transforms in to a string with dots
// just that simple so us from Brazil don't feel excluded
// e.g 834,142.30 -> 834.142,30
func Dot(v float64) string {

	var vWithCommas = Commaf(v)

	// split the number on the dot, "10,012.50" -> "10,012" % ".50"
	var vSplitInComma = strings.Split(vWithCommas, ".")

	var vCommasToDots = strings.Replace(vSplitInComma[0], ",", ".", -1)

	return fmt.Sprintf("%s,%s", vCommasToDots, vSplitInComma[1][0:])
}

// DotWithDigits get the float and put the number of
// decimals as in the argument d
// e.g  v = 834,142.30 and d = 1 -> 834,142.3
func DotWithDigits(v float64, d int) string {
	return stripTrailingDigitsDot(Dot(v), d)
}
