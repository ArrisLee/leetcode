package plusone

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	in := []int{9, 9, 9, 9, 9, 9}
	expect := []int{1, 0, 0, 0, 0, 0, 0}
	out := plusOne(in)
	for i := range out {
		if out[i] != expect[i] {
			t.Error("failed")
		}
	}
}
