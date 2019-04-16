package exist

/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (30.77%)
 * Total Accepted:    267.7K
 * Total Submissions: 868K
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given a 2D board and a word, find if the word exists in the grid.
 *
 * The word can be constructed from letters of sequentially adjacent cell,
 * where "adjacent" cells are those horizontally or vertically neighboring. The
 * same letter cell may not be used more than once.
 *
 * Example:
 *
 *
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 *
 * Given word = "ABCCED", return true.
 * Given word = "SEE", return true.
 * Given word = "ABCB", return false.
 *
 *
 */
func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 {
		return false
	} else if len(word) == 0 {
		return true
	}
	words := []byte(word)
	row := len(board)
	column := len(board[0])
	if len(word) > row*column {
		return false
	}
	return check(board, row, column, words, []pos{})
}

type pos struct {
	x int
	y int
}

func check(board [][]byte, row, column int, target []byte, alreadyIn []pos) bool {
	if len(target) == len(alreadyIn) {
		return true
	}
	if len(alreadyIn) > 0 {
		lastPos := alreadyIn[len(alreadyIn)-1]
		nextChar := target[len(alreadyIn)]
		if lastPos.x-1 >= 0 {
			if nextChar == board[lastPos.x-1][lastPos.y] {
				already := false
				for _, in := range alreadyIn {
					if in.x == lastPos.x-1 && in.y == lastPos.y {
						already = true
						break
					}
				}
				if !already {
					if check(board, row, column, target, append(alreadyIn, pos{lastPos.x - 1, lastPos.y})) {
						return true
					}
				}
			}
		}
		if lastPos.y-1 >= 0 {
			if nextChar == board[lastPos.x][lastPos.y-1] {
				already := false
				for _, in := range alreadyIn {
					if in.x == lastPos.x && in.y == lastPos.y-1 {
						already = true
						break
					}
				}
				if !already {
					if check(board, row, column, target, append(alreadyIn, pos{lastPos.x, lastPos.y - 1})) {
						return true
					}
				}
			}
		}
		if lastPos.y+1 < column {
			if nextChar == board[lastPos.x][lastPos.y+1] {
				already := false
				for _, in := range alreadyIn {
					if in.x == lastPos.x && in.y == lastPos.y+1 {
						already = true
						break
					}
				}
				if !already {
					if check(board, row, column, target, append(alreadyIn, pos{lastPos.x, lastPos.y + 1})) {
						return true
					}
				}
			}
		}
		if lastPos.x+1 < row {
			if nextChar == board[lastPos.x+1][lastPos.y] {
				already := false
				for _, in := range alreadyIn {
					if in.x == lastPos.x+1 && in.y == lastPos.y {
						already = true
						break
					}
				}
				if !already {
					if check(board, row, column, target, append(alreadyIn, pos{lastPos.x + 1, lastPos.y})) {
						return true
					}
				}
			}
		}
		return false
	} else {
		for i := 0; i < row; i++ {
			for j := 0; j < column; j++ {
				if board[i][j] == target[0] {
					if check(board, row, column, target, []pos{pos{i, j}}) {
						return true
					}
				}
			}
		}
		return false
	}
}
