package largestnumber

import "testing"

func TestLargestNumber(t *testing.T) {
	if largestNumber([]int{10, 2}) != "210" {
		t.Error("210 Error")
	}
	if largestNumber([]int{3, 30, 34, 5, 9}) != "9534330" {
		t.Error("9534330 Error")
	}
	if largestNumber([]int{121, 12}) != "12121" {
		t.Error("12121 Error")
	}
	if largestNumber([]int{12, 121}) != "12121" {
		t.Error("12121 Error")
	}
	if largestNumber([]int{44, 444}) != "44444" {
		t.Error("44444 Error")
	}
}

func TestGetNumberPosition(t *testing.T) {
	if getNumberPosition(123, 0) != 1 {
		t.Error("123 0 Error")
	}
	if getNumberPosition(123, 1) != 2 {
		t.Error("123 0 Error")
	}
	if getNumberPosition(123, 2) != 3 {
		t.Error("123 0 Error")
	}
	if getNumberPosition(123, 3) != 1 {
		t.Error("123 0 Error")
	}
}
