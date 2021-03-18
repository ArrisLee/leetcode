package twosums

import "testing"

// func TestTwoSumSltnOne(t *testing.T) {
// 	want := map[int]bool{
// 		0: true,
// 		1: true,
// 	}
// 	got := twoSumSltnOne([]int{2, 7, 11, 15}, 9)
// 	for i := range got {
// 		if !want[got[i]] {
// 			t.Errorf("want: %v, got: %v", want, got)
// 		}
// 	}
// 	got = twoSumSltnOne([]int{3, 2, 4}, 6)
// 	want = map[int]bool{
// 		1: true,
// 		2: true,
// 	}
// 	for i := range got {
// 		if !want[got[i]] {
// 			t.Errorf("want: %v, got: %v", want, got)
// 		}
// 	}
// }

func TestTwoSumSltnTwo(t *testing.T) {
	want := map[int]bool{
		2: true,
		3: true,
	}
	got := twoSum([]int{2, 7, 11, 15}, 26)
	for i := range got {
		if !want[got[i]] {
			t.Errorf("want: %v, got: %v", want, got)
		}
	}
	got = twoSum([]int{3, 2, 4}, 6)
	want = map[int]bool{
		1: true,
		2: true,
	}
	for i := range got {
		if !want[got[i]] {
			t.Errorf("want: %v, got: %v", want, got)
		}
	}
}
