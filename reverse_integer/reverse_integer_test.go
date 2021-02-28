package reverseinteger

import "testing"

func TestReverse(t *testing.T) {
	want := -321
	got := reverse(-123)
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
