package find_bottom_left_tree_value

/*
 * @lc app=leetcode id=513 lang=golang
 *
 * [513] Find Bottom Left Tree Value
 *
 * https://leetcode.com/problems/find-bottom-left-tree-value/description/
 *
 * algorithms
 * Medium (59.69%)
 * Total Accepted:    84.2K
 * Total Submissions: 141K
 * Testcase Example:  '[2,1,3]'
 *
 *
 * Given a binary tree, find the leftmost value in the last row of the tree.
 *
 *
 * Example 1:
 *
 * Input:
 *
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 *
 * Output:
 * 1
 *
 *
 *
 * ⁠ Example 2:
 *
 * Input:
 *
 * ⁠       1
 * ⁠      / \
 * ⁠     2   3
 * ⁠    /   / \
 * ⁠   4   5   6
 * ⁠      /
 * ⁠     7
 *
 * Output:
 * 7
 *
 *
 *
 * Note:
 * You may assume the tree (i.e., the given root node) is not NULL.
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
func findBottomLeftValue(root *TreeNode) int {
	var leftValue int
	nextLevel := make([]*TreeNode, 0)
	currentLevel := make([]*TreeNode, 0)
	currentLevel = append(currentLevel, root)
	for len(currentLevel) > 0 {
		for i, node := range currentLevel {
			if i == 0 {
				leftValue = node.Val
			}
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		currentLevel, nextLevel = nextLevel, currentLevel
		nextLevel = nextLevel[0:0]
	}
	return leftValue
}
