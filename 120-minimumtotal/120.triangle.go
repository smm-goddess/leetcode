package minimumtotal

/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 *
 * https://leetcode.com/problems/triangle/description/
 *
 * algorithms
 * Medium (38.78%)
 * Total Accepted:    176.9K
 * Total Submissions: 454.2K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * Given a triangle, find the minimum path sum from top to bottom. Each step
 * you may move to adjacent numbers on the row below.
 *
 * For example, given the following triangle
 *
 *
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 *
 *
 * The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
 *
 * Note:
 *
 * Bonus point if you are able to do this using only O(n) extra space, where n
 * is the total number of rows in the triangle.
 *
 */
func minimumTotal(triangle [][]int) int {
	if triangle == nil {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	} else {
		min := 0
		for i := len(triangle) - 2; i >= 0; i-- {
			for j := 0; j < len(triangle[i]); j++ {
				min = triangle[i+1][j]
				if triangle[i+1][j+1] < min {
					min = triangle[i+1][j+1]
				}
				triangle[i][j] = triangle[i][j] + min
			}
		}
	}
	return triangle[0][0]
}
