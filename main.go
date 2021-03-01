package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	convertNode(807)
}

func convertNode(num int) *ListNode {
	result := &ListNode{
		Val: num % 10,
	}
	current := result
	num /= 10
	for num > 0 {
		if current.Next == nil {
			current.Next = &ListNode{
				Val: num % 10,
			}
		}
		current = current.Next
		num /= 10
	}
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
	return result
}
