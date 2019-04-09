package combinationsum2

import "sort"

/*
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

Each number in candidates may only be used once in the combination.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8,
A solution set is:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
Example 2:

Input: candidates = [2,5,2,1,2], target = 5,
A solution set is:
[
  [1,2,2],
  [5]
]
*/
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	retVal := make([][]int, 0, 16)
	combinationSumIter(candidates, target, []int{}, &retVal)
	return retVal
}

func combinationSumIter(candidates []int, target int, alreadyIn []int, retVal *[][]int) {
	if target == 0 {
		*retVal = append(*retVal, append([]int{}, alreadyIn...))
		return
	} else if target > 0 && len(candidates) > 0 {
		combinationSumIter(candidates[1:], target-candidates[0], append(alreadyIn, candidates[0]), retVal)
		sameIndex := 1
		for sameIndex < len(candidates) && candidates[sameIndex] == candidates[sameIndex-1] {
			sameIndex++
		}
		combinationSumIter(candidates[sameIndex:], target, alreadyIn, retVal)
	}
}
