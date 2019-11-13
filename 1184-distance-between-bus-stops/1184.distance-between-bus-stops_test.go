package distance_between_bus_stops

import "testing"

func TestDistance(t *testing.T) {
	distance := []int{1, 2, 3, 4}
	if distanceBetweenBusStops(distance, 0, 1) != 1 {
		t.Error("1 error")
	}
	if distanceBetweenBusStops(distance, 0, 2) != 3 {
		t.Error("3 error")
	}
	if distanceBetweenBusStops(distance, 0, 3) != 4 {
		t.Error("4 error")
	}
	if distanceBetweenBusStops(distance, 0, 0) != 0 {
		t.Error("0 error")
	}
}
