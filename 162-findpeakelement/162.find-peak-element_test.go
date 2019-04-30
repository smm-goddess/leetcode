package findpeakelement

import "testing"

func TestFindPeakElement(t *testing.T) {
	if findPeakElement([]int{1, 2, 3, 1}) != 2 {
		t.Error("1 Error")
	}
	if findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}) != 1 {
		t.Error("2 Error")
	}
}
