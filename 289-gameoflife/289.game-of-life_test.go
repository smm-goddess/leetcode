package gemeoflive

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

func listOfListEquals(a, b [][]int) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil || len(a) != len(b) {
		return false
	} else {
		for i, v := range a {
			if !listEquals(b[i], v) {
				return false
			}
		}
		return true
	}
}

func TestGameOfLive(t *testing.T) {
	startState := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	nextState := [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
	}
	gameOfLife(startState)
	if !listOfListEquals(startState, nextState) {
		t.Error("Test Fail")
	}
}

//func TestCountAlive(t *testing.T) {
//	startState := [][]int{
//		{1, 0, 1},
//		{0, 0, 1},
//		{1, 1, 1},
//		{0, 0, 0},
//	}
//	rows = 4
//	columns = 3
//	for i := 0; i < 4; i++ {
//		for j := 0; j < 3; j++ {
//			fmt.Println("(", i, ",", j, ")", " ", countAlive(startState, i, j))
//		}
//	}
//}
