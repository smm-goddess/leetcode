package minimum_absolute_difference_in_bst

import "math"

/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
 *
 * https://leetcode.com/problems/minimum-absolute-difference-in-bst/description/
 *
 * algorithms
 * Easy (51.76%)
 * Total Accepted:    70.7K
 * Total Submissions: 136.7K
 * Testcase Example:  '[1,null,3,2]'
 *
 * Given a binary search tree with non-negative values, find the minimum
 * absolute difference between values of any two nodes.
 *
 * Example:
 *
 *
 * Input:
 *
 * ⁠  1
 * ⁠   \
 * ⁠    3
 * ⁠   /
 * ⁠  2
 *
 * Output:
 * 1
 *
 * Explanation:
 * The minimum absolute difference is 1, which is the difference between 2 and
 * 1 (or between 2 and 3).
 *
 *
 *
 *
 * Note: There are at least two nodes in this BST.
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func min(a ...int) int {
	m := math.MaxInt32
	for _, d := range a {
		if d < m {
			m = d
		}
	}
	return m
}

func distance(a, b *TreeNode) int {
	if a == nil || b == nil {
		return math.MaxInt32
	} else {
		dis := a.Val - b.Val
		if dis < 0 {
			dis = -dis
		}
		return dis
	}
}

func getMinimumDifference(root *TreeNode) int {
	if root == nil || (root.Right == nil && root.Left == nil) {
		return math.MaxInt32
	} else {
		return min(
			distance(root, rightMost(root.Left)),
			distance(root, leftMost(root.Right)),
			getMinimumDifference(root.Left),
			getMinimumDifference(root.Right))
	}
}

func rightMost(root *TreeNode) *TreeNode {
	right := root
	for right != nil && right.Right != nil {
		right = right.Right
	}
	return right
}
func leftMost(root *TreeNode) *TreeNode {
	left := root
	for left != nil && left.Left != nil {
		left = left.Left
	}
	return left
}
