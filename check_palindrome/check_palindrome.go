package checkpalindrome

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	str := strconv.Itoa(x)
	return checkStr(str)
}

func checkStr(in string) bool {
	endOne := len(in) / 2
	endTwo := endOne
	if len(in)%2 != 0 {
		endTwo = endOne + 1
	}
	for i, j := 0, len(in)-1; i < endOne && j >= endTwo; {
		if in[i] != in[j] {
			return false
		}
		i++
		j--
	}
	return true
}
