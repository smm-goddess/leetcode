package out_of_boundary_paths

import (
	"fmt"
	"testing"
)

var testCases = [][]int{
	{2, 2, 2, 0, 0},
	{1, 3, 3, 0, 1},
	{8, 7, 16, 1, 5},
}
var testCasesAnswers = []int{
	6,
	12,
	102984580,
}

func Test_findPaths(t *testing.T) {
	for index, testCase := range testCases {
		m, n, N, i, j := testCase[0], testCase[1], testCase[2], testCase[3], testCase[4]
		if ans := findPaths(m, n, N, i, j); ans != testCasesAnswers[index] {
			t.Error(fmt.Sprintf("%d wrong,ans:%d", index, ans))
		}
	}
}
