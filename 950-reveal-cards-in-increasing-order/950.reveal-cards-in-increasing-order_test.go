package deck_revealed_increasing

import "testing"

func listEquals(a, b []int) bool {
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

func TestDeckRevealedIncreasing(t *testing.T) {
	input := []int{17, 13, 11, 2, 3, 5, 7}
	output := []int{2, 13, 3, 11, 5, 17, 7}
	if !listEquals(deckRevealedIncreasing(input), output) {
		t.Error("error")
	}
	if !listEquals(deckRevealedIncreasing([]int{1}), []int{1}) {
		t.Error("1 error")
	}
}
