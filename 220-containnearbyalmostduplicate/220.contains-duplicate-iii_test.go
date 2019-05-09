package containsnearbyalmostduplicate

import (
	"fmt"
	"testing"
)

type testCase struct {
	array  []int
	k      int
	t      int
	answer bool
}

func (tc *testCase) String() string {
	return fmt.Sprintf("%d,%d ", tc.k, tc.t)
}

var testCases = []*testCase{
	{
		[]int{1, 2, 3, 1},
		3,
		0,
		true,
	},
	{
		[]int{1, 0, 1, 1},
		1,
		2,
		true,
	},
	{
		[]int{1, 5, 9, 1, 5, 9},
		2,
		3,
		false,
	},
}

func TestContainsDuplicate(t *testing.T) {
	for _, tc := range testCases {
		if containsNearbyAlmostDuplicate(tc.array, tc.k, tc.t) != tc.answer {
			t.Error(fmt.Sprintf("%s Error", tc))
		}
	}
}
