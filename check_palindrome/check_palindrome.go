package checkpalindrome

func isPalindrome(x int) bool {
	// always return false for negative numer
	if x < 0 {
		return false
	}
	// normal case: check if x eqauls to its reversed number
	reversed := 0
	for tmp := x; tmp != 0; tmp /= 10 {
		reversed = reversed*10 + tmp%10
	}
	return x == reversed
}

// func isPalindrome(x int) bool {
// 	if x < 0 {
// 		return false
// 	}
// 	if x < 10 {
// 		return true
// 	}
// 	str := strconv.Itoa(x)
// 	return checkStr(str)
// }

// func checkStr(in string) bool {
// 	endOne := len(in) / 2
// 	endTwo := endOne
// 	if len(in)%2 != 0 {
// 		endTwo = endOne + 1
// 	}
// 	for i, j := 0, len(in)-1; i < endOne && j >= endTwo; {
// 		if in[i] != in[j] {
// 			return false
// 		}
// 		i++
// 		j--
// 	}
// 	return true
// }
