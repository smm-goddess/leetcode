package graycode

import (
	"fmt"
	"testing"
)

var grayCodeTestSample = map[int][]int{
	//0: {0},
	// 	1: {0, 1},
	2: {0, 1, 3, 2},
	3: {0, 1, 3, 2, 6, 7, 5, 4}}

func sliceEquals(a, b []int) bool {
	if a == nil && b == nil {
		return true
	} else if len(a) != len(b) {
		return false
	} else {
		for i, v := range a {
			if b[i] != v {
				return false
			}
		}
		return true
	}
}

func TestGrayCode(t *testing.T) {
	for k, v := range grayCodeTestSample {
		if !sliceEquals(grayCode(k), v) {
			t.Error(fmt.Sprintf("%d gray code Error", k))
		}
	}
}
