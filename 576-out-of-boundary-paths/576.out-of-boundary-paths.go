package out_of_boundary_paths

/*
 * @lc app=leetcode id=576 lang=golang
 *
 * [576] Out of Boundary Paths
 *
 * https://leetcode.com/problems/out-of-boundary-paths/description/
 *
 * algorithms
 * Medium (31.60%)
 * Total Accepted:    20.5K
 * Total Submissions: 62.2K
 * Testcase Example:  '2\n2\n2\n0\n0'
 *
 * There is an m by n grid with a ball. Given the start coordinate (i,j) of the
 * ball, you can move the ball to adjacent cell or cross the grid boundary in
 * four directions (up, down, left, right). However, you can at most move N
 * times. Find out the number of paths to move the ball out of grid boundary.
 * The answer may be very large, return it after mod 10^9 + 7.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: m = 2, n = 2, N = 2, i = 0, j = 0
 * Output: 6
 * Explanation:
 *
 *
 *
 * Example 2:
 *
 *
 * Input: m = 1, n = 3, N = 3, i = 0, j = 1
 * Output: 12
 * Explanation:
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * Once you move the ball out of boundary, you cannot move it back.
 * The length and height of the grid is in range [1,50].
 * N is in range [0,50].
 *
 *
 */
func findPaths(m int, n int, N int, i int, j int) int {
	matrix := make([][]int, m, m)
	cacheMatrix := make([][]int, m, m)
	for index := 0; index < m; index++ {
		matrix[index] = make([]int, n, n)
		cacheMatrix[index] = make([]int, n, n)
	}
	for ; N > 0; N-- {
		for row := 0; row < m; row++ {
			for column := 0; column < n; column++ {
				cacheMatrix[row][column] = (matrixValue(matrix, m, n, row, column-1) +
					matrixValue(matrix, m, n, row, column+1) +
					matrixValue(matrix, m, n, row-1, column) +
					matrixValue(matrix, m, n, row+1, column)) % 1000000007
			}
		}
		matrix, cacheMatrix = cacheMatrix, matrix
	}
	return matrix[i][j]
}

func matrixValue(matrix [][]int, m, n, row, column int) int {
	if row < 0 || row >= m || column < 0 || column >= n {
		return 1
	} else {
		return matrix[row][column]
	}
}
