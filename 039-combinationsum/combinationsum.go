package combinationsum

import (
	"sort"
)

/*
[39] Combination Sum

https://leetcode.com/problems/combination-sum/description/

* algorithms
* Medium (47.42%)
* Total Accepted:    319.9K
* Total Submissions: 674.7K
* Testcase Example:  '[2,3,6,7]\n7'

Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

The same repeated number may be chosen from candidates unlimited number of times.

Note:


        All numbers (including target) will be positive integers.
        The solution set must not contain duplicate combinations.


Example 1:


Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]


Example 2:


Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/

func combinationSumRecur(candidates []int, target int) [][]int {
	retVal := make([][]int, 0)
	for _, v := range candidates {
		newCandi := make([]int, 0)
		newTarget := target - v
		for _, value := range candidates {
			if value >= v && value <= newTarget {
				newCandi = append(newCandi, value)
			}
		}
		if len(newCandi) == 0 {
			if newTarget == 0 {
				retVal = append(retVal, []int{v})
			}
		} else {
			for _, can := range combinationSumRecur(newCandi, newTarget) {
				retVal = append(retVal, append([]int{v}, can...))
			}
		}
	}
	return retVal
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	retVal := make([][]int, 0)
	combinationIter(candidates, 0, target, []int{}, &retVal)
	return retVal
}

func combinationIter(candidates []int, index, target int, path []int, retVal *[][]int) {
	if target == 0 {
		p := make([]int, len(path))
		copy(p, path)
		*retVal = append(*retVal, p)
	} else if target > 0 {
		for i, v := range candidates {
			if i >= index && v <= target {
				combinationIter(candidates, i, target-v, append(path, v), retVal)
			}
		}
	}
}
