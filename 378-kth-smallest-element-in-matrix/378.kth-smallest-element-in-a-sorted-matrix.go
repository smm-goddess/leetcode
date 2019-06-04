package smallest

import (
	"math"
)

/*
 * @lc app=leetcode id=378 lang=golang
 *
 * [378] Kth Smallest Element in a Sorted Matrix
 *
 * https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/description/
 *
 * algorithms
 * Medium (48.89%)
 * Total Accepted:    109.6K
 * Total Submissions: 221.5K
 * Testcase Example:  '[[1,5,9],[10,11,13],[12,13,15]]\n8'
 *
 * Given a n x n matrix where each of the rows and columns are sorted in
 * ascending order, find the kth smallest element in the matrix.
 *
 *
 * Note that it is the kth smallest element in the sorted order, not the kth
 * distinct element.
 *
 *
 * Example:
 *
 * matrix = [
 * ⁠  [ 1,  5,  9],
 * ⁠  [10, 11, 13],
 * ⁠  [12, 13, 15]
 * ],
 * k = 8,
 *
 * return 13.
 *
 *
 *
 * Note:
 * You may assume k is always valid, 1 ≤ k ≤ n2.
 */
func kthSmallest(matrix [][]int, k int) int {
	size := len(matrix)
	indexes := make([]int, size, size)
	smallestIndex, smallestValue := 0, 0
	for k > 0 {
		smallestValue = math.MaxInt32
		for i := 0; i < size; i++ {
			if indexes[i] < size {
				if matrix[i][indexes[i]] < smallestValue {
					smallestValue = matrix[i][indexes[i]]
					smallestIndex = i
				}
			}
		}
		indexes[smallestIndex]++
		k--
	}
	return smallestValue
}
