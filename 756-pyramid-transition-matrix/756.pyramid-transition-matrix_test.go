package pyramid_transition

import (
	"testing"
)

type TestCase struct {
	bottom  string
	allowed []string
	answer  bool
}

var testCases []*TestCase

func init() {
	testCases = []*TestCase{
		{
			"BCD",
			[]string{"BCG", "CDE", "GEA", "FFF"},
			true,
		},
		{
			"AABA",
			[]string{"AAA", "AAB", "ABA", "ABB", "BAC"},
			false,
		},
	}
}

func TestPyramidTransition(t *testing.T) {
	for _, testCase := range testCases {
		if testCase.answer != pyramidTransition(testCase.bottom, testCase.allowed) {
			t.Error(testCase.bottom, " error")
		}
	}
}

//func TestFindAvailable(t *testing.T) {
//	fmt.Println(findAvailable([]string{"AAA", "AAB", "ABA", "ABB", "BAC"}, "BA"))
//}
