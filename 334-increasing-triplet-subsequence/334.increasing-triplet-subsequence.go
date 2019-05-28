package increasingtriplet

import "math"

/*
 * @lc app=leetcode id=334 lang=golang
 *
 * [334] Increasing Triplet Subsequence
 *
 * https://leetcode.com/problems/increasing-triplet-subsequence/description/
 *
 * algorithms
 * Medium (39.48%)
 * Total Accepted:    92.1K
 * Total Submissions: 232.8K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Given an unsorted array return whether an increasing subsequence of length 3
 * exists or not in the array.
 *
 * Formally the function should:
 *
 * Return true if there exists i, j, k
 * such that arr[i] < arr[j] < arr[k] given 0 ≤ i < j < k ≤ n-1 else return
 * false.
 *
 * Note: Your algorithm should run in O(n) time complexity and O(1) space
 * complexity.
 *
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,4,5]
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [5,4,3,2,1]
 * Output: false
 *
 *
 *
 */
func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt32, math.MaxInt32
	for _, v := range nums {
		if v <= first {
			first = v
		} else if v <= second {
			second = v
		} else {
			return true
		}
	}
	return false
}
