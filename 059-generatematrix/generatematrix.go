package generatematrix

/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 *
 * https://leetcode.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (45.88%)
 * Total Accepted:    131.6K
 * Total Submissions: 286.7K
 * Testcase Example:  '3'
 *
 * Given a positive integer n, generate a square matrix filled with elements
 * from 1 to n^2 in spiral order.
 *
 * Example:
 *
 *
 * Input: 3
 * Output:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 8, 9, 4 ],
 * ⁠[ 7, 6, 5 ]
 * ]
 *
 *
 */
func generateMatrix(n int) [][]int {
	x, y, directX, directY := 0, 0, 0, 1
	answer := make([][]int, n)
	for i := 0; i < n; i++ {
		answer[i] = make([]int, n)
	}
	for i := 1; i <= n*n; i++ {
		answer[x][y] = i
		nextPosition(&x, &y, &directX, &directY, n)
	}
	return answer
}

func nextPosition(x, y, directX, directY *int, n int) {
	if *directX == 0 && *directY == 1 && *x+*y == n-1 {
		*directX, *directY = 1, 0
	} else if *directX == 1 && *directY == 0 && *x == *y {
		*directX, *directY = 0, -1
	} else if *directX == 0 && *directY == -1 && *x+*y == n-1 {
		*directX, *directY = -1, 0
	} else if *directX == -1 && *directY == 0 && *x-*y == 1 {
		*directX, *directY = 0, 1
	}
	*x, *y = *x+*directX, *y+*directY
}
