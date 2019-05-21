package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	if singleNumber([]int{2, 2, 1}) != 1 {
		t.Error("1 error")
	}
	if singleNumber([]int{4, 1, 2, 1, 2}) != 4 {
		t.Error("4 error")
	}
}
