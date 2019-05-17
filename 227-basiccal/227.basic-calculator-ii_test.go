package calculator

import "testing"

func TestCalculator(t *testing.T) {
	if calculate("3 + 2 * 2") != 7 {
		t.Error("3 + 2 * 2 Error")
	}
	if calculate("3 / 2") != 1 {
		t.Error("3 / 2 Error")
	}
	if calculate("3 + 5 / 2") != 5 {
		t.Error("3 + 5 / 2 Error")
	}
	if calculate("321 * 53 + 12 - 123/4") != 16995 {
		t.Error("321 * 53 + 12 - 123/4")
	}
	if calculate(" 3/2 ") != 1 {
		t.Error("3/2 Error")
	}
	if calculate(" 2/3 ") != 0 {
		t.Error("2/3 Error")
	}
	if calculate("1-1+1*5") != 5 {
		t.Error("1-1+1 Error")
	}
}
