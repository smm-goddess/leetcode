package countdigits

import "testing"

func TestCountDigits(t *testing.T) {
	if countNumbersWithUniqueDigits(2) != 91 {
		t.Error("2 error")
	}
	if countNumbersWithUniqueDigits(1) != 10 {
		t.Error("1 error")
	}
}
