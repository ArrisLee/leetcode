package findsamenumbers

import "testing"

func TestFindSameNumbers(t *testing.T) {
	in := []int{2, 1, 3, 1, 4}
	if out, err := findSameNumber(in); out != 1 || err != nil {
		t.Errorf("wrong")
	}
}
