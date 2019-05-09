package rob

/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 *
 * https://leetcode.com/problems/house-robber-ii/description/
 *
 * algorithms
 * Medium (35.17%)
 * Total Accepted:    113.4K
 * Total Submissions: 322K
 * Testcase Example:  '[2,3,2]'
 *
 * You are a professional robber planning to rob houses along a street. Each
 * house has a certain amount of money stashed. All houses at this place are
 * arranged in a circle. That means the first house is the neighbor of the last
 * one. Meanwhile, adjacent houses have security system connected and it will
 * automatically contact the police if two adjacent houses were broken into on
 * the same night.
 *
 * Given a list of non-negative integers representing the amount of money of
 * each house, determine the maximum amount of money you can rob tonight
 * without alerting the police.
 *
 * Example 1:
 *
 *
 * Input: [2,3,2]
 * Output: 3
 * Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money
 * = 2),
 * because they are adjacent houses.
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money =
 * 3).
 * Total amount you can rob = 1 + 3 = 4.
 *
 */
func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else {
		length := len(nums)
		sum := make([]int, length)
		sum[0] = nums[0]
		for i := 1; i < length-1; i++ {
			for j := 0; j <= i; j++ {
				if j >= 2 {
					sum[i] = max(sum[i], sum[j-2]+nums[j])
				} else {
					sum[i] = max(sum[i], nums[j])
				}
			}
		}
		r1 := sum[length-2]

		sum[1] = nums[1]
		for i := 2; i < length; i++ {
			sum[i] = 0
			for j := 1; j <= i; j++ {
				if j >= 3 {
					sum[i] = max(sum[i], sum[j-2]+nums[j])
				} else {
					sum[i] = max(sum[i], nums[j])
				}
			}
		}
		r2 := sum[length-1]
		return max(r1, r2)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
