package matrix

/*
 * @lc app=leetcode id=542 lang=golang
 *
 * [542] 01 Matrix
 *
 * https://leetcode.com/problems/01-matrix/description/
 *
 * algorithms
 * Medium (35.03%)
 * Total Accepted:    42.1K
 * Total Submissions: 118.6K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * Given a matrix consists of 0 and 1, find the distance of the nearest 0 for
 * each cell.
 *
 * The distance between two adjacent cells is 1.
 *
 *
 *
 * Example 1:
 *
 *
 * Input:
 * [[0,0,0],
 * ⁠[0,1,0],
 * ⁠[0,0,0]]
 *
 * Output:
 * [[0,0,0],
 * [0,1,0],
 * [0,0,0]]
 *
 *
 * Example 2:
 *
 *
 * Input:
 * [[0,0,0],
 * ⁠[0,1,0],
 * ⁠[1,1,1]]
 *
 * Output:
 * [[0,0,0],
 * ⁠[0,1,0],
 * ⁠[1,2,1]]
 *
 *
 *
 *
 * Note:
 *
 *
 * The number of elements of the given matrix will not exceed 10,000.
 * There are at least one 0 in the given matrix.
 * The cells are adjacent in only four directions: up, down, left and right.
 *
 *
 */
var MAX = 20000

func updateMatrix(matrix [][]int) [][]int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	} else {
		rows, columns := len(matrix), len(matrix[0])
		dist := make([][]int, rows, rows)
		for i := range dist {
			dist[i] = make([]int, columns)
		}
		for i := 0; i < rows; i++ {
			for j := 0; j < columns; j++ {
				if matrix[i][j] == 0 {
					dist[i][j] = 0
				} else {
					dist[i][j] = MAX
					if i > 0 {
						dist[i][j] = min(dist[i][j], dist[i-1][j]+1)
					}
					if j > 0 {
						dist[i][j] = min(dist[i][j], dist[i][j-1]+1)
					}
				}
			}
		}
		for i := rows - 1; i >= 0; i-- {
			for j := columns - 1; j >= 0; j-- {
				if i < rows-1 {
					dist[i][j] = min(dist[i][j], dist[i+1][j]+1)
				}
				if j < columns-1 {
					dist[i][j] = min(dist[i][j], dist[i][j+1]+1)
				}
			}
		}
		return dist
	}
}

func min(a ...int) (min int) {
	min = MAX
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return
}
