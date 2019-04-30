package fractiontodecimal

import "testing"

func TestFractionToDecimal(t *testing.T) {
	if fractionToDecimal(1, 2) != "0.5" {
		t.Error("1/2 Error")
	}
	if fractionToDecimal(2, 1) != "2" {
		t.Error("2/1 Error")
	}
	if fractionToDecimal(2, 3) != "0.(6)" {
		t.Error("2/3 Error")
	}
	if fractionToDecimal(8, 7) != "1.(142857)" {
		t.Error("8/7 Error")
	}
	if fractionToDecimal(0, 3) != "0" {
		t.Error("0/3 Error")
	}
}
