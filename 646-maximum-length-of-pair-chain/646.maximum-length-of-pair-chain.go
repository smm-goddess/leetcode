package find_longest_chain

import (
	"sort"
)

/*
 * @lc app=leetcode id=646 lang=golang
 *
 * [646] Maximum Length of Pair Chain
 *
 * https://leetcode.com/problems/maximum-length-of-pair-chain/description/
 *
 * algorithms
 * Medium (49.96%)
 * Total Accepted:    43.6K
 * Total Submissions: 87.3K
 * Testcase Example:  '[[1,2], [2,3], [3,4]]'
 *
 *
 * You are given n pairs of numbers. In every pair, the first number is always
 * smaller than the second number.
 *
 *
 *
 * Now, we define a pair (c, d) can follow another pair (a, b) if and only if b
 * < c. Chain of pairs can be formed in this fashion.
 *
 *
 *
 * Given a set of pairs, find the length longest chain which can be formed. You
 * needn't use up all the given pairs. You can select pairs in any order.
 *
 *
 *
 * Example 1:
 *
 * Input: [[1,2], [2,3], [3,4]]
 * Output: 2
 * Explanation: The longest chain is [1,2] -> [3,4]
 *
 *
 *
 * Note:
 *
 * The number of given pairs will be in the range [1, 1000].
 *
 *
 */

type ps [][]int

func (this ps) Len() int {
	return len(this)
}

func (this ps) Less(i, j int) bool {
	return this[i][0] < this[j][0]
}

func (this ps) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func findLongestChain(pairs [][]int) int {
	sort.Sort(ps(pairs))
	distance := make([]int, len(pairs), len(pairs))
	for i := 0; i < len(pairs); i++ {
		distance[i] = 1
		for j := 0; j < i; j++ {
			if pairs[j][1] < pairs[i][0] {
				distance[i] = max(distance[i], distance[j]+1)
			}
		}
	}
	ans := distance[0]
	for i := 0; i < len(pairs); i++ {
		if ans < distance[i] {
			ans = distance[i]
		}
	}
	return ans
}
