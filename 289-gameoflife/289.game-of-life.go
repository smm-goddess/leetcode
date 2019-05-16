package gemeoflive

/*
 * @lc app=leetcode id=289 lang=golang
 *
 * [289] Game of Life
 *
 * https://leetcode.com/problems/game-of-life/description/
 *
 * algorithms
 * Medium (44.28%)
 * Total Accepted:    110.9K
 * Total Submissions: 246.1K
 * Testcase Example:  '[[0,1,0],[0,0,1],[1,1,1],[0,0,0]]'
 *
 * According to the Wikipedia's article: "The Game of Life, also known simply
 * as Life, is a cellular automaton devised by the British mathematician John
 * Horton Conway in 1970."
 *
 * Given a board with m by n cells, each cell has an initial state live (1) or
 * dead (0). Each cell interacts with its eight neighbors (horizontal,
 * vertical, diagonal) using the following four rules (taken from the above
 * Wikipedia article):
 *
 *
 * Any live cell with fewer than two live neighbors dies, as if caused by
 * under-population.
 * Any live cell with two or three live neighbors lives on to the next
 * generation.
 * Any live cell with more than three live neighbors dies, as if by
 * over-population..
 * Any dead cell with exactly three live neighbors becomes a live cell, as if
 * by reproduction.
 *
 *
 * Write a function to compute the next state (after one update) of the board
 * given its current state. The next state is created by applying the above
 * rules simultaneously to every cell in the current state, where births and
 * deaths occur simultaneously.
 *
 * Example:
 *
 *
 * Input:
 * [
 * [0,1,0],
 * [0,0,1],
 * [1,1,1],
 * [0,0,0]
 * ]
 * Output:
 * [
 * [0,0,0],
 * [1,0,1],
 * [0,1,1],
 * [0,1,0]
 * ]
 *
 *
 * Follow up:
 *
 *
 * Could you solve it in-place? Remember that the board needs to be updated at
 * the same time: You cannot update some cells first and then use their updated
 * values to update other cells.
 * In this question, we represent the board using a 2D array. In principle, the
 * board is infinite, which would cause problems when the active area
 * encroaches the border of the array. How would you address these problems?
 *
 *
 */
var diff = []int{-1, 0, 1}
var rows, columns int

func gameOfLife(board [][]int) {
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return
	}
	rows, columns = len(board), len(board[0])
	stashedUpdate, stashedUpdateNext := make([]int, columns, columns), make([]int, columns, columns)
	var alive int
	var updateRow []int
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if i >= 2 {
				board[i-2][j] = stashedUpdate[j]
			}
			alive = countAlive(board, i, j)
			updateRow = stashedUpdate
			if i == 1 {
				updateRow = stashedUpdateNext
			}
			if board[i][j] == 0 {
				if alive == 3 {
					updateRow[j] = 1
				} else {
					updateRow[j] = 0
				}
			} else {
				if alive == 2 || alive == 3 {
					updateRow[j] = 1
				} else {
					updateRow[j] = 0
				}
			}
		}
		if i >= 2 {
			stashedUpdate, stashedUpdateNext = stashedUpdateNext, stashedUpdate
		}
	}
	if rows == 1 {
		for j := 0; j < columns; j++ {
			board[0][j] = stashedUpdate[j]
		}
	} else {
		for j := 0; j < columns; j++ {
			board[rows-2][j] = stashedUpdate[j]
			board[rows-1][j] = stashedUpdateNext[j]
		}
	}
}

func countAlive(board [][]int, row, column int) (alive int) {
	var x, y int
	for _, dx := range diff {
		for _, dy := range diff {
			if !(dx == 0 && dy == 0) {
				x, y = row+dx, column+dy
				if x >= 0 && x < rows && y >= 0 && y < columns {
					if board[x][y] == 1 {
						alive++
					}
				}
			}
		}
	}
	return
}
