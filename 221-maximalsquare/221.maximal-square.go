package maxsquare

/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 *
 * https://leetcode.com/problems/maximal-square/description/
 *
 * algorithms
 * Medium (32.56%)
 * Total Accepted:    127.7K
 * Total Submissions: 388.4K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * Given a 2D binary matrix filled with 0's and 1's, find the largest square
 * containing only 1's and return its area.
 *
 * Example:
 *
 *
 * Input:
 *
 * 1 0 1 0 0
 * 1 0 1 1 1
 * 1 1 1 1 1
 * 1 0 0 1 0
 *
 * Output: 4
 *
 */
var rows, columns int

func maximalSquare(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, columns = len(matrix), len(matrix[0])
	maxSideLength, sideLength := 0, 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == '1' {
				sideLength = 1
				for sidesAreSquare(matrix, i, j, sideLength+1) {
					sideLength = sideLength + 1
				}
				if sideLength > maxSideLength {
					maxSideLength = sideLength
				}
			}
		}
	}
	return maxSideLength * maxSideLength
}

func sidesAreSquare(matric [][]byte, row, column, sideLength int) bool {
	if row+sideLength > rows || column+sideLength > columns {
		return false
	}
	for i := 0; i < sideLength; i++ {
		if matric[row+sideLength-1][column+i] != '1' || matric[row+i][column+sideLength-1] != '1' {
			return false
		}
	}
	return true
}
