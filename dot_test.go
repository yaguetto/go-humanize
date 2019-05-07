package humanize

import (
	"log"
	"strings"
	"testing"
)

func TestDot(t *testing.T) {

	var tests = []string{"134,012.50", "109,876,542,210.50", "-553,835,342,934.50", "0.00"}

	for i := range tests {

		var vSplitInComma = strings.Split(tests[i], ".")

		var vCommasToDots = strings.Replace(vSplitInComma[0], ",", ".", -1)

		log.Printf("%s,%s", vCommasToDots, vSplitInComma[1][0:])
	}

}
