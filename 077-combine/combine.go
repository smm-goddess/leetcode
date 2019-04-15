package combine

/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 *
 * https://leetcode.com/problems/combinations/description/
 *
 * algorithms
 * Medium (46.74%)
 * Total Accepted:    193.8K
 * Total Submissions: 414K
 * Testcase Example:  '4\n2'
 *
 * Given two integers n and k, return all possible combinations of k numbers
 * out of 1 ... n.
 *
 * Example:
 *
 *
 * Input: n = 4, k = 2
 * Output:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 *
 */
func combine(n int, k int) [][]int {
	if k == 0 || k > n {
		return [][]int{}
	} else {
		var result [][]int
		candidates := make([]int, n)
		for i := 1; i <= n; i++ {
			candidates[i-1] = i
		}
		resultIter(candidates, []int{}, k, &result)
		return result
	}
}

func resultIter(candidates, alreadyIn []int, mustSelect int, result *[][]int) {
	if len(alreadyIn) == mustSelect {
		*result = append(*result, append([]int{}, alreadyIn...))
	} else if mustSelect > len(alreadyIn) {
		if len(candidates) == mustSelect-len(alreadyIn) {
			*result = append(*result, append(append([]int{}, alreadyIn...), candidates...))
		} else if len(candidates) > mustSelect-len(alreadyIn) {
			resultIter(candidates[1:], alreadyIn, mustSelect, result)
			resultIter(candidates[1:], append(alreadyIn, candidates[0]), mustSelect, result)
		}
	}
}
