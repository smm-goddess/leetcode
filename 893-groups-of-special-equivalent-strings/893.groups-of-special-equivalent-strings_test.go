package num_special_equiv_groups

import "testing"

type Case struct {
	A     []string
	Count int
}

var cases []*Case

func init() {
	cases = []*Case{
		{
			[]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"},
			3,
		},
		{
			[]string{"abc", "acb", "bac", "bca", "cab", "cba"},
			3,
		},
	}
}

func TestNumSpecialEquivGroups(t *testing.T) {
	for _, c := range cases {
		if numSpecialEquivGroups(c.A) != c.Count {
			t.Error(c.A, " error")
		}
	}
}
