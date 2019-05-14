package maxsquare

import "testing"

func TestMaxSquare(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	if maximalSquare(matrix) != 4 {
		t.Error("4 Error")
	}
	matrix = [][]byte{
		{'0', '0', '0', '1'},
		{'1', '1', '0', '1'},
		{'1', '1', '1', '1'},
		{'0', '1', '1', '1'},
		{'0', '1', '1', '1'}}
	if maximalSquare(matrix) != 9 {
		t.Error("9 Error")
	}
	matrix = [][]byte{{'1'}}
	if maximalSquare(matrix) != 1 {
		t.Error("1 Error")
	}
}
