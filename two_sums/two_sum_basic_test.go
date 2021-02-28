package twosums

import "testing"

func TestTwoSum(t *testing.T) {
	want := []int{0, 1}
	got := twoSum([]int{2, 7, 11, 15}, 9)
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("want: %v, got: %v", want, got)
	}

	want = []int{1, 2}
	got = twoSum([]int{3, 2, 4}, 6)
	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
