package lis

import "testing"

func TestLongestIncreasingSubsequence(t *testing.T) {
	if lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}) != 4 {
		t.Error("Error")
	}
	if lengthOfLIS([]int{10, 9, 2, 5, 3, 4}) != 3 {
		t.Error("Error")
	}
}
