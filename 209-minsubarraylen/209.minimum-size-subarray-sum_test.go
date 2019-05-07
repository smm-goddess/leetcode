package minsubarraylen

import "testing"

func TestMinSubarrayLen(t *testing.T) {
	if minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}) != 2 {
		t.Error("7 Error")
	}
	if minSubArrayLen(4, []int{1, 4, 4}) != 1 {
		t.Error("4 Error")
	}
	if minSubArrayLen(11, []int{1, 2, 3, 4, 5}) != 3 {
		t.Error("11 Error")
	}
}
