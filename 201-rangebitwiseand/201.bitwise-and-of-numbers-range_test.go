package rangebitwiseand

import "testing"

func TestRangeBitwiseAnd(t *testing.T) {
	if rangeBitwiseAnd(5, 7) != 4 {
		t.Error("5 7 Error")
	}
	if rangeBitwiseAnd(0, 1) != 0 {
		t.Error("0 1 Error")
	}
}
