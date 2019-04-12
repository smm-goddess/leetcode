package minpathsum

/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 *
 * https://leetcode.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (46.16%)
 * Total Accepted:    220K
 * Total Submissions: 475.9K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * Given a m x n grid filled with non-negative numbers, find a path from top
 * left to bottom right which minimizes the sum of all numbers along its path.
 *
 * Note: You can only move either down or right at any point in time.
 *
 * Example:
 *
 *
 * Input:
 * [
 * [1,3,1],
 * ⁠ [1,5,1],
 * ⁠ [4,2,1]
 * ]
 * Output: 7
 * Explanation: Because the path 1→3→1→1→1 minimizes the sum.
 *
 *
 */
func minPathSum(grid [][]int) int {
	rows := len(grid)
	columns := len(grid[0])
	board := make([][]int, rows)
	for i := 0; i < rows; i++ {
		board[i] = make([]int, columns)
	}
	for i := rows - 1; i >= 0; i-- {
		for j := columns - 1; j >= 0; j-- {
			if i+1 < rows && j+1 < columns {
				if board[i+1][j] < board[i][j+1] {
					board[i][j] = grid[i][j] + board[i+1][j]
				} else {
					board[i][j] = grid[i][j] + board[i][j+1]
				}
			} else if i+1 < rows {
				board[i][j] = grid[i][j] + board[i+1][j]
			} else if j+1 < columns {
				board[i][j] = grid[i][j] + board[i][j+1]
			} else {
				board[i][j] = grid[i][j]
			}
		}
	}
	return board[0][0]
}
