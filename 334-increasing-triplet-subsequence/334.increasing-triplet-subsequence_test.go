package increasingtriplet

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	if !increasingTriplet([]int{1, 2, 3, 4, 5}) {
		t.Error("error")
	}
	if increasingTriplet([]int{5, 4, 3, 2, 1}) {
		t.Error("error")
	}
}
