package combinationSum3

/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 *
 * https://leetcode.com/problems/combination-sum-iii/description/
 *
 * algorithms
 * Medium (50.91%)
 * Total Accepted:    119.8K
 * Total Submissions: 233.6K
 * Testcase Example:  '3\n7'
 *
 *
 * Find all possible combinations of k numbers that add up to a number n, given
 * that only numbers from 1 to 9 can be used and each combination should be a
 * unique set of numbers.
 *
 * Note:
 *
 *
 * All numbers will be positive integers.
 * The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: k = 3, n = 7
 * Output: [[1,2,4]]
 *
 *
 * Example 2:
 *
 *
 * Input: k = 3, n = 9
 * Output: [[1,2,6], [1,3,5], [2,3,4]]
 *
 *
 */
func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	candidates := make([]int, 9, 9)
	for i := range candidates {
		candidates[i] = i + 1
	}
	combinationSum(k, n, []int{}, candidates, &result)
	return result
}

func combinationSum(numCount, target int, alreadyIn, candidates []int, result *[][]int) {
	if numCount == 0 && target == 0 {
		*result = append(*result, alreadyIn)
	} else if len(candidates) == 0 || numCount*candidates[len(candidates)-1] < target {
		return
	} else {
		max := candidates[len(candidates)-1]
		if max <= target {
			combinationSum(numCount, target, alreadyIn, candidates[:len(candidates)-1], result)
			combinationSum(numCount-1, target-max, append([]int{max}, alreadyIn...), candidates[:len(candidates)-1], result)
		} else {
			combinationSum(numCount, target, alreadyIn, candidates[:target], result)
		}
	}
}
