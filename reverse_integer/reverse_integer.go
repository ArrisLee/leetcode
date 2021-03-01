package reverseinteger

import (
	"math"
	"strconv"
)

func reverseIntSltnOne(x int) int {
	orig := strconv.Itoa(x)
	numStr := ""
	end := 0
	if string(orig[0]) == "-" {
		numStr = "-"
		end = 1
	}
	for i := len(orig) - 1; i >= end; i-- {
		numStr += string(orig[i])
	}
	if result, _ := strconv.Atoi(numStr); result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	} else {
		return result
	}
}

func reverseIntSltnTwo(x int) int {
	reversed := 0
	for x != 0 {
		reversed *= 10
		reversed += x % 10
		x /= 10
	}
	return reversed
}
