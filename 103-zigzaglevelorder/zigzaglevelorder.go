package zigzaglevelorder

/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (40.57%)
 * Total Accepted:    198K
 * Total Submissions: 488K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the zigzag level order traversal of its nodes'
 * values. (ie, from left to right, then right to left for the next level and
 * alternate between).
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
 * return its zigzag level order traversal as:
 *
 * [
 * ⁠ [3],
 * ⁠ [20,9],
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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	preSlice := []*TreeNode{root}
	var result [][]int
	var afterSlice []*TreeNode
	right := true
	for len(preSlice) > 0 {
		var level []int
		for i := len(preSlice) - 1; i >= 0; i-- {
			node := preSlice[i]
			level = append(level, node.Val)
			if right {
				if node.Left != nil {
					afterSlice = append(afterSlice, node.Left)
				}
				if node.Right != nil {
					afterSlice = append(afterSlice, node.Right)
				}
			} else {
				if node.Right != nil {
					afterSlice = append(afterSlice, node.Right)
				}
				if node.Left != nil {
					afterSlice = append(afterSlice, node.Left)
				}
			}
		}
		right = !right
		result = append(result, level)
		preSlice, afterSlice = afterSlice, preSlice
		afterSlice = afterSlice[:0]
	}
	return result
}
