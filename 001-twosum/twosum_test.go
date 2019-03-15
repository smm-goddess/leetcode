package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	a := []int{2, 7, 11, 15}
	r := twoSum(a, 9)
	if r[0] != 0 || r[1] != 1 {
		t.Error("fail")
	}
}
