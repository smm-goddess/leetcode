package measurewater

import "testing"

func TestCanMeasureWater(t *testing.T) {
	if !canMeasureWater(3, 5, 4) {
		t.Error("3 5 4")
	}
	if canMeasureWater(2, 6, 5) {
		t.Error("2 6 5")
	}
}
