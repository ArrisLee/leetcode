package main

import (
	"fmt"
)

func main() {
	check := []int{1, 2, 3, 4, 5}
	one := check[1:]
	two := check[1:3]
	three := check[:3]

	fmt.Println(one, two, three)
}
