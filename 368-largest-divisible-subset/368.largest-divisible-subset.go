package largestdivisiblesubset

import (
	"sort"
)

/*
 * @lc app=leetcode id=368 lang=golang
 *
 * [368] Largest Divisible Subset
 *
 * https://leetcode.com/problems/largest-divisible-subset/description/
 *
 * algorithms
 * Medium (34.61%)
 * Total Accepted:    46.4K
 * Total Submissions: 133.4K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a set of distinct positive integers, find the largest subset such that
 * every pair (Si, Sj) of elements in this subset satisfies:
 *
 * Si % Sj = 0 or Sj % Si = 0.
 *
 * If there are multiple solutions, return any subset is fine.
 *
 * Example 1:
 *
 *
 *
 * Input: [1,2,3]
 * Output: [1,2] (of course, [1,3] will also be ok)
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,4,8]
 * Output: [1,2,4,8]
 *
 *
 *
 */
func largestDivisibleSubset(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	length := len(nums)
	max, parent := make([]int, length, length), make([]int, length, length)
	for i := range parent {
		parent[i] = -1
		max[i] = 1
	}
	max[0] = 1
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && max[i] < max[j]+1 {
				max[i], parent[i] = max[j]+1, j
			}
		}
	}
	maxValue, maxIndex := 0, 0
	for i, m := range max {
		if m > maxValue {
			maxIndex = i
			maxValue = m
		}
	}
	var ans []int
	for maxIndex >= 0 {
		ans = append(ans, nums[maxIndex])
		maxIndex = parent[maxIndex]
	}
	return ans
}
