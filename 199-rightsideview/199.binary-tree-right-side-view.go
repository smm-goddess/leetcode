package rightsideview

/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
 *
 * https://leetcode.com/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (47.20%)
 * Total Accepted:    163.4K
 * Total Submissions: 343.4K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * Given a binary tree, imagine yourself standing on the right side of it,
 * return the values of the nodes you can see ordered from top to bottom.
 * 
 * Example:
 * 
 * 
 * Input: [1,2,3,null,5,null,4]
 * Output: [1, 3, 4]
 * Explanation:
 * 
 * ⁠  1            <---
 * ⁠/   \
 * 2     3         <---
 * ⁠\     \
 * ⁠ 5     4       <---
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	thisLevel := []*TreeNode{root}
	var nextLevel []*TreeNode
	var rightSideView []int
	for len(thisLevel) > 0 {
		rightSideView = append(rightSideView, thisLevel[len(thisLevel)-1].Val)
		for i := 0; i < len(thisLevel); i++ {
			if thisLevel[i].Left != nil {
				nextLevel = append(nextLevel, thisLevel[i].Left)
			}
			if thisLevel[i].Right != nil {
				nextLevel = append(nextLevel, thisLevel[i].Right)
			}
		}
		thisLevel = nextLevel
		nextLevel = make([]*TreeNode, 0)
	}
	return rightSideView
}
