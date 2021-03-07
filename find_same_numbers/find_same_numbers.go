package findsamenumbers

import "fmt"

func findSameNumber(numbers []int) (int, error) {
	for i := range numbers {
		if numbers[i] != i {
			temp := numbers[i]
			numbers[i] = i
			if numbers[temp] == temp {
				return numbers[temp], nil
			}
			numbers[temp] = temp
		}
	}
	return 0, fmt.Errorf("not identified")
}
