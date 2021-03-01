package addtwonumbers

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	want := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	got := addTwoNumbers(l1, l2)
	for want.Next != nil {
		if want.Val != got.Val {
			t.Errorf("want: %d, got: %d\n", want.Val, got.Val)
		}
		want = want.Next
		got = got.Next
	}
}
