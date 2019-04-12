package searchmatrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50}}
	//if !searchMatrix(matrix, 3) {
	//	t.Error("Search 3 Error")
	//}
	//if searchMatrix(matrix, 13) {
	//	t.Error("Search 13 Error")
	//}
	if !searchMatrix(matrix, 30) {
		t.Error("Search 30 Error")
	}
	//if !searchMatrix([][]int{{1}}, 1) {
	//	t.Error("Search 1 Error")
	//}
	//if !searchMatrix([][]int{{1}, {3}}, 3) {
	//	t.Error("Search 1 Error")
	//}
}

//
//func TestSearchIndex(t *testing.T) {
//	starts := []int{1, 10, 23}
//	if searchIndex(starts, 5) != 0 {
//		t.Error("Search Index Error")
//	}
//	if searchIndex(starts, 10) != 1 {
//		t.Error("Search Index Error")
//	}
//	if searchIndex(starts, 11) != 1 {
//		t.Error("Search Index Error")
//	}
//
//	if searchIndex(starts, 23) != 2 {
//		t.Error("Search Index Error")
//	}
//}
