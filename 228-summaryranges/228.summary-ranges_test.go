package summaryranges

import "testing"

func stringSliceEquals(a, b []string) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil || len(a) != len(b) {
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

func TestSummaryRanges(t *testing.T) {
	if !stringSliceEquals(summaryRanges([]int{0, 1, 2, 4, 5, 7}), []string{"0->2", "4->5", "7"}) {
		t.Error("0->7 Error")
	}
	if !stringSliceEquals(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}), []string{"0", "2->4", "6", "8->9"}) {
		t.Error("0->9 Error")
	}
}
