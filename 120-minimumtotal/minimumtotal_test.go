package minimumtotal

import "testing"

func TestMinimumTotal(t *testing.T) {
	triagle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	if minimumTotal(triagle) != 11 {
		t.Error("Test Fail")
	}
}
