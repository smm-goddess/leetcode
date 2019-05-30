package integerbreak

import "testing"

func TestIntegerBreak(t *testing.T) {
	if integerBreak(2) != 1 {
		t.Error("2 error")
	}
	if integerBreak(10) != 36 {
		t.Error("10 error")
	}
}
