package fliplight

import "testing"

type Case struct {
	N      int
	M      int
	Output int
}

var cases = []Case{
	{1, 1, 2,},
	{2, 1, 3,},
	{3, 1, 4},
	{3, 2, 7},
}

func TestFlipLight(t *testing.T) {
	for _, cs := range cases {
		if cs.Output != flipLights(cs.N, cs.M) {
			t.Error(cs.N, cs.M, cs.Output, " error")
		}
	}
}
