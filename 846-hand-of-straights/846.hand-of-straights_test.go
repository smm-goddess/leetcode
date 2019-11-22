package is_n_straight_hand

import "testing"

type Case struct {
	Hand   []int
	W      int
	Output bool
}

var cases = []*Case{
	{
		[]int{1, 2, 3, 6, 2, 3, 4, 7, 8},
		3,
		true,
	},
	{
		[]int{1, 2, 3, 4, 5},
		4,
		false,
	},
}

func TestIsNStraightHand(t *testing.T) {
	for _, cs := range cases {
		if cs.Output != isNStraightHand(cs.Hand, cs.W) {
			t.Error(cs.W, " error")
		}
	}
}
