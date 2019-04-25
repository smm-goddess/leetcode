package gasstation

import "testing"

func TestCanCompleteCurcuit(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	if canCompleteCircuit(gas, cost) != 3 {
		t.Error("First Fail")
	}
	gas = []int{2, 3, 4}
	cost = []int{3, 4, 5}
	if canCompleteCircuit(gas, cost) != -1 {
		t.Error("Secone Fail")
	}
}
