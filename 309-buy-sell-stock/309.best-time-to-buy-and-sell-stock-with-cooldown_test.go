package maxprofit

import "testing"

func TestMaxProfit(t *testing.T) {
	if maxProfit([]int{1, 2, 3, 0, 2}) != 3 {
		t.Error("error")
	}
}
