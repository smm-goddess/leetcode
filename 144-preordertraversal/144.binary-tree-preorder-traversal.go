package preordertraversal

/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Medium (50.78%)
 * Total Accepted:    327.8K
 * Total Submissions: 642.4K
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the preorder traversal of its nodes' values.
 * 
 * Example:
 * 
 * 
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 * 
 * Output: [1,2,3]
 * 
 * 
 * Follow up: Recursive solution is trivial, could you do it iteratively?
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
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var array []int
	nodes := []*TreeNode{root}
	preorderTraversalIter(&array, nodes)
	return array
}

func preorderTraversalIter(array *[]int, nodes []*TreeNode) {
	if nodes == nil || len(nodes) == 0 {
		return
	} else {
		first := nodes[len(nodes)-1]
		remains := nodes[:len(nodes)-1]
		*array = append(*array, first.Val)
		if first.Right != nil {
			remains = append(remains, first.Right)
		}
		if first.Left != nil {
			remains = append(remains, first.Left)
		}
		preorderTraversalIter(array, remains)
	}
}
