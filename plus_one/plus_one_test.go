package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	in := []int{9, 9, 9, 9, 9, 9}
	expect := []int{1, 0, 0, 0, 0, 0, 0}
	got := plusOne(in)
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("want %v, got: %v", expect, got)
	}
}
