package numtrees

/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 *
 * https://leetcode.com/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (45.59%)
 * Total Accepted:    193.2K
 * Total Submissions: 422.4K
 * Testcase Example:  '3'
 *
 * Given n, how many structurally unique BST's (binary search trees) that store
 * values 1 ... n?
 *
 * Example:
 *
 *
 * Input: 3
 * Output: 5
 * Explanation:
 * Given n = 3, there are a total of 5 unique BST's:
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 *
 *
 */

var nums = make(map[int]int)

func numTrees(n int) int {
	switch n {
	case 0, 1:
		return 1
	case 2:
		return 2
	case 3:
		return 5
	default:
		if v, ok := nums[n]; ok {
			return v
		} else {
			num := 0
			for i := 1; i <= n; i++ {
				num += numTrees(i-1) * numTrees(n-i)
			}
			nums[n] = num
			return num
		}
	}
}
