package checkpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	table := map[int]bool{
		0:     true,
		6:     true,
		1001:  true,
		99999: true,
		1101:  false,
		8686:  false,
		-1111: false,
	}
	for in, want := range table {
		if got := isPalindrome(in); got != want {
			t.Errorf("current case: %d, want: %v, got: %v", in, want, got)
		}
	}
}
