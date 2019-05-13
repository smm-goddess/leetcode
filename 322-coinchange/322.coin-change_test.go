package coinchange

import "testing"

func TestCoinChange(t *testing.T) {
	if coinChange([]int{1}, 0) != 0 {
		t.Error("0 Error")
	}
	if coinChange([]int{1}, 1) != 1 {
		t.Error("1 Error")
	}
	if coinChange([]int{1, 2, 5}, 11) != 3 {
		t.Error("11 Error")
	}
	if coinChange([]int{2}, 3) != -1 {
		t.Error("3 Error")
	}
	if coinChange([]int{3, 7, 405, 436}, 8839) != 25 {
		t.Error("8839 Error")
	}
}
