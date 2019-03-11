package validbst

/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 *
 * https://leetcode.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (25.28%)
 * Total Accepted:    363.9K
 * Total Submissions: 1.4M
 * Testcase Example:  '[2,1,3]'
 *
 * Given a binary tree, determine if it is a valid binary search tree (BST).
 *
 * Assume a BST is defined as follows:
 *
 *
 * The left subtree of a node contains only nodes with keys less than the
 * node's key.
 * The right subtree of a node contains only nodes with keys greater than the
 * node's key.
 * Both the left and right subtrees must also be binary search trees.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 * Output: false
 * Explanation: The input is: [5,1,4,null,null,3,6]. The root node's
 * value
 * is 5 but its right child's value is 4.
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

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	right, left := root.Right, root.Left
	if right == nil && left == nil {
		return true
	} else if right == nil {
		valid := left.Val < root.Val && isValidBST(left)
		if !valid {
			return false
		} else {
			max := left
			for max.Right != nil {
				max = max.Right
			}
			return root.Val > max.Val
		}
	} else if left == nil {
		valid := right.Val > root.Val && isValidBST(right)
		if !valid {
			return false
		} else {
			min := right
			for min.Left != nil {
				min = min.Left
			}
			return root.Val < min.Val
		}
	} else {
		valid := right.Val > root.Val && left.Val < root.Val && isValidBST(right) && isValidBST(left)
		if !valid {
			return false
		} else {
			max := left
			for max.Right != nil {
				max = max.Right
			}
			if root.Val > max.Val {
				min := right
				for min.Left != nil {
					min = min.Left
				}
				return root.Val < min.Val
			} else {
				return false
			}
		}
	}
}
