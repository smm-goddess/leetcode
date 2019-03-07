package main

import "fmt"

func main() {
	board := make([][]byte, 9, 9)
	board[0] = []byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'}
	board[1] = []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'}
	board[2] = []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'}
	board[3] = []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'}
	board[4] = []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'}
	board[5] = []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'}
	board[6] = []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'}
	board[7] = []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'}
	board[8] = []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	checkValid := func(line []byte) bool {
		count := []bool{false, false, false, false, false, false, false, false, false}
		for _, v := range line {
			if v != '.' {
				pos := v - '1'
				if count[pos] {
					return false
				} else {
					count[pos] = true
				}
			}
		}
		return true
	}

	for i := 0; i < 9; i ++ {
		row := make([]byte, 9, 9)
		col := make([]byte, 9, 9)
		square := make([]byte, 9, 9)
		for j := 0; j < 9; j++ {
			row[j] = board[i][j]
			col[j] = board[j][i]
			a, b, c, d := i/3, i%3, j/3, j%3
			square[j] = board[a*3+c][b*3+d]
		}
		if !(checkValid(row) && checkValid(col) && checkValid(square)) {
			return false
		}
	}
	return true
}

