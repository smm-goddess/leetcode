package longest_univalue_path

/*
 * @lc app=leetcode id=687 lang=golang
 *
 * [687] Longest Univalue Path
 *
 * https://leetcode.com/problems/longest-univalue-path/description/
 *
 * algorithms
 * Easy (33.47%)
 * Total Accepted:    71.1K
 * Total Submissions: 205K
 * Testcase Example:  '[5,4,5,1,1,5]'
 *
 * Given a binary tree, find the length of the longest path where each node in
 * the path has the same value. This path may or may not pass through the
 * root.
 *
 * The length of path between two nodes is represented by the number of edges
 * between them.
 *
 *
 *
 * Example 1:
 *
 * Input:
 *
 *
 * ⁠             5
 * ⁠            / \
 * ⁠           4   5
 * ⁠          / \   \
 * ⁠         1   1   5
 *
 *
 * Output: 2
 *
 *
 *
 * Example 2:
 *
 * Input:
 *
 *
 * ⁠             1
 * ⁠            / \
 * ⁠           4   5
 * ⁠          / \   \
 * ⁠         4   4   5
 *
 *
 * Output: 2
 *
 *
 *
 * Note: The given binary tree has not more than 10000 nodes. The height of the
 * tree is not more than 1000.
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(length ...int) int {
	m := 0
	for _, v := range length {
		if m < v {
			m = v
		}
	}
	return m
}
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(longest(0, root.Val, root.Left)+
		longest(0, root.Val, root.Right),
		longestUnivaluePath(root.Right),
		longestUnivaluePath(root.Left))
}

func longest(length, preValue int, root *TreeNode) int {
	if root == nil || root.Val != preValue {
		return length
	} else {
		return max(longest(length+1, preValue, root.Right), longest(length+1, preValue, root.Left))
	}
}
