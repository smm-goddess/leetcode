package minsubarraylen

/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 *
 * https://leetcode.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (34.50%)
 * Total Accepted:    173.4K
 * Total Submissions: 499.6K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * Given an array of n positive integers and a positive integer s, find the
 * minimal length of a contiguous subarray of which the sum ≥ s. If there isn't
 * one, return 0 instead.
 * 
 * Example: 
 * 
 * 
 * Input: s = 7, nums = [2,3,1,2,4,3]
 * Output: 2
 * Explanation: the subarray [4,3] has the minimal length under the problem
 * constraint.
 * 
 * Follow up:
 * 
 * If you have figured out the O(n) solution, try coding another solution of
 * which the time complexity is O(n log n). 
 * 
 */
func minSubArrayLen(s int, nums []int) int {
	minLen, sum, j := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		sum, j = nums[i], i+1
		if sum >= s {
			return 1
		}
		for sum < s && j < len(nums) {
			if minLen > 0 && j-i >= minLen {
				break
			}
			sum += nums[j]
			j++
			if sum >= s {
				minLen = j - i
			}
		}
	}
	return minLen
}
