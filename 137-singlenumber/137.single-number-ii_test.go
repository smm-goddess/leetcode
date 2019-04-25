package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	array := []int{2, 2, 3, 2}
	if singleNumber(array) != 3 {
		t.Error("First Fail")
	}
	array = []int{0, 1, 0, 1, 0, 1, 99}
	if singleNumber(array) != 99 {
		t.Error("Secone Fail")
	}
}
