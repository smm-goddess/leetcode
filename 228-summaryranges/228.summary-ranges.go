package summaryranges

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 *
 * https://leetcode.com/problems/summary-ranges/description/
 *
 * algorithms
 * Medium (35.52%)
 * Total Accepted:    130K
 * Total Submissions: 363.1K
 * Testcase Example:  '[0,1,2,4,5,7]'
 *
 * Given a sorted integer array without duplicates, return the summary of its
 * ranges.
 *
 * Example 1:
 *
 *
 * Input:  [0,1,2,4,5,7]
 * Output: ["0->2","4->5","7"]
 * Explanation: 0,1,2 form a continuous range; 4,5 form a continuous range.
 *
 *
 * Example 2:
 *
 *
 * Input:  [0,2,3,4,6,8,9]
 * Output: ["0","2->4","6","8->9"]
 * Explanation: 2,3,4 form a continuous range; 8,9 form a continuous range.
 *
 *
 */
func summaryRanges(nums []int) []string {
	r := make([]string, 0)
	summaryIter(nums, &r)
	return r
}

func summaryIter(nums []int, result *[]string) {
	if nums == nil || len(nums) == 0 {
		return
	} else if len(nums) == 1 {
		*result = append(*result, strconv.Itoa(nums[0]))
		return
	} else {
		start, continueIndex := nums[0], 1
		for continueIndex < len(nums) && nums[continueIndex]-start == 1 {
			start = nums[continueIndex]
			continueIndex++
		}
		if nums[0] == start {
			*result = append(*result, strconv.Itoa(start))
		} else {
			*result = append(*result, fmt.Sprintf("%d->%d", nums[0], start))
		}
		summaryIter(nums[continueIndex:], result)
	}
}
