package sortedlisttobst
/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 *
 * https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/description/
 *
 * algorithms
 * Medium (40.31%)
 * Total Accepted:    171.2K
 * Total Submissions: 424.8K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * Given a singly linked list where elements are sorted in ascending order,
 * convert it to a height balanced BST.
 *
 * For this problem, a height-balanced binary tree is defined as a binary tree
 * in which the depth of the two subtrees of every node never differ by more
 * than 1.
 *
 * Example:
 *
 *
 * Given the sorted linked list: [-10,-3,0,5,9],
 *
 * One possible answer is: [0,-3,9,-10,null,5], which represents the following
 * height balanced BST:
 *
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	i, j, k := head, head.Next, head
	if j == nil {
		return &TreeNode{i.Val, nil, nil}
	} else {
		for j != nil {
			i = i.Next
			if j.Next != nil {
				j = j.Next
			} else {
				j = nil
			}
			if j != nil && j.Next != nil {
				j = j.Next
			} else {
				j = nil
			}
		}
		for k.Next.Val != i.Val {
			k = k.Next
		}
		k.Next = nil
		return &TreeNode{i.Val, sortedListToBST(head), sortedListToBST(i.Next)}
	}
}
