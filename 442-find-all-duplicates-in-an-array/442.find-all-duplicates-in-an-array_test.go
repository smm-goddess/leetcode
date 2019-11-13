package find_all_duplicates_in_an_array

import (
	"sort"
	"testing"
)

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

func TestFindDuplicates(t *testing.T) {
	//if !listEquals(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}), []int{2, 3}) {
	//	t.Error("error")
	//}
	expected := []int{14, 8, 6, 29, 5, 46, 39, 47, 18, 17, 36, 10, 24, 11, 25}
	sort.Ints(expected)
	if !listEquals(findDuplicates([]int{39, 31, 8, 14, 14, 38, 5, 15, 29, 49, 18, 6, 30, 47, 8, 35, 2, 17, 6, 10, 29, 46, 41, 48, 1, 36, 5, 28, 46, 39, 7, 47, 18, 42, 17, 11, 36, 45, 21, 33, 24, 10, 24, 50, 25, 16, 9, 12, 11, 25}),
		expected) {
		t.Error("long error")
	}
}
