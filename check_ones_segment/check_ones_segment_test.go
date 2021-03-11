package checkonessegment

import "testing"

func TestCheckOnesSegment(t *testing.T) {
	table := map[string]bool{
		"1":                true,
		"10":               true,
		"101":              false,
		"11111":            true,
		"1111100000111111": false,
	}
	for ts, want := range table {
		if got := checkOnesSegment(ts); got != want {
			t.Errorf("current test case: %s, want: %t, got: %t", ts, want, got)
		}
	}
}
