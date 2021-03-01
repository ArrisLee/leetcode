package reverseinteger

import "testing"

func TestReverseIntSltnOne(t *testing.T) {
	want := -321
	got := reverseIntSltnOne(-123)
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestReverseIntSltnTwo(t *testing.T) {
	want := -321
	got := reverseIntSltnTwo(-123)
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
	want = 10000000001
	got = reverseIntSltnTwo(10000000001)
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
	want = 0
	got = reverseIntSltnTwo(0)
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
