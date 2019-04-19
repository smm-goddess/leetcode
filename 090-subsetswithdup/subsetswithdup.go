package subsetswithdup

import (
	"sort"
)

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 *
 * https://leetcode.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (41.86%)
 * Total Accepted:    196.8K
 * Total Submissions: 468.6K
 * Testcase Example:  '[1,2,2]'
 *
 * Given a collection of integers that might contain duplicates, nums, return
 * all possible subsets (the power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 *
 * Input: [1,2,2]
 * Output:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 *
 *
 */
func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	subsetsIter(&result, []int{}, nums)
	return result
}

func setsEquals(a, b []int) bool {
	if (a == nil && b == nil) || (len(a) != len(b)) {
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

func subsetsIter(result *[][]int, alreadyIn, remains []int) {
	if remains == nil || len(remains) == 0 {
		already := false
		in := append([]int{}, alreadyIn...)
		sort.Ints(in)
		for _, v := range *result {
			if setsEquals(v, in) {
				already = true
			}
		}
		if !already {
			*result = append(*result, in)
		}
	} else {
		subsetsIter(result, alreadyIn, remains[1:])
		subsetsIter(result, append(alreadyIn, remains[0]), remains[1:])
	}
}
