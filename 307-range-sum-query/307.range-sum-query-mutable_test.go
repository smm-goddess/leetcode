package rangesumquery

import "testing"

func TestRangeSumQuery(t *testing.T) {
	numArray := Constructor([]int{1, 3, 5})
	if numArray.SumRange(0, 2) != 9 {
		t.Error("9 error")
	}
	numArray.Update(0, 2)
	if numArray.SumRange(0, 2) != 10 {
		t.Error("8 error")
	}
}
