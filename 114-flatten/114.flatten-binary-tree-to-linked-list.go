package flatten

// import "fmt"

/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 *
 * https://leetcode.com/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (41.68%)
 * Total Accepted:    230.6K
 * Total Submissions: 550.6K
 * Testcase Example:  '[1,2,5,3,4,null,6]'
 *
 * Given a binary tree, flatten it to a linked list in-place.
 *
 * For example, given the following tree:
 *
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   5
 * ⁠/ \   \
 * 3   4   6
 *
 *
 * The flattened tree should look like:
 *
 *
 * 1
 * ⁠\
 * ⁠ 2
 * ⁠  \
 * ⁠   3
 * ⁠    \
 * ⁠     4
 * ⁠      \
 * ⁠       5
 * ⁠        \
 * ⁠         6
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
func flatten(root *TreeNode) {
	if root != nil {
		preorder(root)
	}
}

func preorder(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	} else {
		start, end := root, root
		leftStart, leftEnd := preorder(start.Left)
		rightStart, rightEnd := preorder(start.Right)
		if leftStart == nil {
			if rightStart != nil {
				start.Left = nil
				start.Right = rightStart
				end = rightEnd
			}
		} else {
			start.Left = nil
			start.Right = leftStart
			if rightStart != nil {
				leftEnd.Right = rightStart
				end = rightEnd
			} else {
				end = leftEnd
			}
		}
		return start, end
	}
}
