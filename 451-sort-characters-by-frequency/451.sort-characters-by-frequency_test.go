package frequency_sort

import "testing"

type Case struct {
	input  string
	output string
}

var cases []*Case

func init() {
	cases = []*Case{
		{
			"tree",
			"eetr",
		},
		{
			"cccaaa",
			"cccaaa",
		},
		{
			"Aabb",
			"bbAa",
		},
		{
			"",
			"",
		},
		{
			"a",
			"a",
		},
	}
}

func TestFrequencySort(t *testing.T) {
	for _, cs := range cases {
		if frequencySort(cs.input) != cs.output {
			t.Error(cs.input, " error")
		}
	}
}
