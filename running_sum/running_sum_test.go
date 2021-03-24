package runningsum

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	expect := []int{1, 3, 6, 10}
	if out := runningSum([]int{1, 2, 3, 4}); !reflect.DeepEqual(out, expect) {
		t.Errorf("failed, want: %v, got: %v", expect, out)
	}
}
