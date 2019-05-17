package majorityelement

/*
 * @lc app=leetcode id=229 lang=golang
 *
 * [229] Majority Element II
 *
 * https://leetcode.com/problems/majority-element-ii/description/
 *
 * algorithms
 * Medium (31.61%)
 * Total Accepted:    99.6K
 * Total Submissions: 311.8K
 * Testcase Example:  '[3,2,3]'
 *
 * Given an integer array of size n, find all elements that appear more than ⌊
 * n/3 ⌋ times.
 *
 * Note: The algorithm should run in linear time and in O(1) space.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3]
 * Output: [3]
 *
 * Example 2:
 *
 *
 * Input: [1,1,1,3,3,2,2,2]
 * Output: [1,2]
 *
 */
func majorityElement(nums []int) []int {
	numberCount := make(map[int]int)
	var ans []int
	minCount := len(nums) / 3
	for _, n := range nums {
		if v, ok := numberCount[n]; ok {
			numberCount[n] = v + 1
		} else {
			numberCount[n] = 1
		}
		if numberCount[n] == minCount+1 {
			ans = append(ans, n)
		}
	}
	return ans
}
