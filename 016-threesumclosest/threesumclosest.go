package threesumclosest

import (
	"sort"
)

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 *
 * https://leetcode.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (41.21%)
 * Total Accepted:    306.5K
 * Total Submissions: 725.1K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * Given an array nums of n integers and an integer target, find three integers
 * in nums such that the sum is closest to target. Return the sum of the three
 * integers. You may assume that each input would have exactly one solution.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 2, 1, -4], and target = 1.
 *
 * The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 *
 *
 */

func distance(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[length-1]
	for i := 0; i < length-2; i++ {
		start := i + 1
		end := length - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum == target {
				return sum
			} else if sum > target {
				end--
			} else {
				start++
			}
			if distance(result, target) > distance(sum, target) {
				result = sum
			}
		}
	}
	return result
}
