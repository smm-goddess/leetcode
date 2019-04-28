package solve

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'O', 'O'},
		{'X', 'X', 'O', 'X'},
		{'X', 'X', 'X', 'X'}}
	solve(board)
	fmt.Println(board)
}
