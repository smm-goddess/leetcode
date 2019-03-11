package levelorder

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (47.17%)
 * Total Accepted:    341.3K
 * Total Submissions: 723.4K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the level order traversal of its nodes' values.
 * (ie, from left to right, level by level).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its level order traversal as:
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
 *
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
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	preSlice := []*TreeNode{root}
	var result [][]int
	var afterSlice []*TreeNode
	for len(preSlice) > 0 {
		var level []int
		for _, node := range preSlice {
			level = append(level, node.Val)
			if node.Left != nil {
				afterSlice = append(afterSlice, node.Left)
			}
			if node.Right != nil {
				afterSlice = append(afterSlice, node.Right)
			}
		}
		result = append(result, level)
		preSlice, afterSlice = afterSlice, preSlice
		afterSlice = afterSlice[:0]
	}
	return result
}
