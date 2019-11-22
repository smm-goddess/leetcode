package largest_time_from_digits

import "testing"

type Case struct {
	Input  []int
	Output string
}

var cases = []Case{
	{
		[]int{1, 2, 3, 4},
		"23:41",
	},
	{
		[]int{5, 5, 5, 5},
		"",
	},
	{
		[]int{0, 0, 0, 0},
		"00:00",
	},
	{
		[]int{9, 5, 3, 2},
		"23:59",
	},
	{
		[]int{0, 5, 0, 0},
		"05:00",
	},
}

func TestLargestTime(t *testing.T) {
	for _, cs := range cases {
		if cs.Output != largestTimeFromDigits(cs.Input) {
			t.Error(cs.Input, " error")
		}
	}
}
