package reverseinteger

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	orig := strconv.Itoa(x)
	numStr := ""
	for i := len(orig) - 1; i >= 0; i-- {
		if i == 0 && string(orig[i]) == "-" {
			numStr = "-" + numStr
		} else {
			numStr += string(orig[i])
		}
	}
	if result, _ := strconv.Atoi(numStr); result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	} else {
		return result
	}
}
