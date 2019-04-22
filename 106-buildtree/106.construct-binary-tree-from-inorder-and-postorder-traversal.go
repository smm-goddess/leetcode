package buildtree
/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (38.76%)
 * Total Accepted:    148.2K
 * Total Submissions: 382.3K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * Given inorder and postorder traversal of a tree, construct the binary tree.
 *
 * Note:
 * You may assume that duplicates do not exist in the tree.
 *
 * For example, given
 *
 *
 * inorder = [9,3,15,20,7]
 * postorder = [9,15,7,20,3]
 *
 * Return the following binary tree:
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if inorder == nil || len(inorder) == 0 {
		return nil
	} else {
		rootVal := postorder[len(postorder)-1]
		root := &TreeNode{rootVal, nil, nil}
		index := 0
		for ; index < len(inorder); index++ {
			if inorder[index] == rootVal {
				break
			}
		}
		root.Left = buildTree(inorder[:index], postorder[:index])
		root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
		return root
	}
}
