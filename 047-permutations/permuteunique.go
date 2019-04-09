package permuteunique

import "sort"

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 *
 * https://leetcode.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (39.74%)
 * Total Accepted:    231.2K
 * Total Submissions: 581.5K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a collection of numbers that might contain duplicates, return all
 * possible unique permutations.
 *
 * Example:
 *
 *
 * Input: [1,1,2]
 * Output:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 *
 *
 */

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	origin := append([]int{}, nums...)
	answer := [][]int{origin}
	nextPermutation(nums)
	for !equal(nums, origin) {
		answer = append(answer, append([]int{}, nums...))
		nextPermutation(nums)
	}
	return answer
}

func equal(a, b []int) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil || len(a) != len(b) {
		return false
	} else {
		for i, v := range a {
			if b[i] != v {
				return false
			}
		}
		return true
	}
}

func nextPermutation(nums []int) {
	length := len(nums)
	i := length - 1
	for ; i > 0; i-- {
		if nums[i-1] < nums[i] {
			i--
			break
		}
	}
	start := i
	for j := length - 1; j > 0; j-- {
		if nums[j] > nums[i] {
			start = i + 1
			nums[i], nums[j] = nums[j], nums[i]
			break
		}
	}
	end := length - 1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
