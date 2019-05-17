package producearray

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
func TestProduceArray(t *testing.T) {
	if !listEquals(productExceptSelf([]int{1, 2, 3, 4}), []int{24, 12, 8, 6}) {
		t.Error("Error")
	}
}
