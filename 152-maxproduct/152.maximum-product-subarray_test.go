package maxproduct

import "testing"

func TestMaxProduct(t *testing.T) {
	if maxProduct([]int{2, 3, -2, 4}) != 6 {
		t.Error("6 Error")
	}
	if maxProduct([]int{-2, 0, -1}) != 0 {
		t.Error("0 Error")
	}
	if maxProduct([]int{-4, -3, -2}) != 12 {
		t.Error("12 Error")
	}
	if maxProduct([]int{0, -2, 0}) != 0 {
		t.Error("12 Error")
	}
}
