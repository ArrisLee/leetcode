package validateparentheses

import "testing"

func TestIsValid(t *testing.T) {
	table := map[string]bool{
		"{[]()}": true,
		"{}[]()": true,
		"{[}":    false,
		")":      false,
		"(":      false,
		"":       true,
	}
	for k, v := range table {
		if check := isValid(k); check != v {
			t.Errorf("current test case: %s, want: %v, got: %v", k, v, check)
		}
	}
}
