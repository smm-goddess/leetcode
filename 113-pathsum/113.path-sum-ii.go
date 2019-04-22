package pathsum

// import "fmt"

/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 *
 * https://leetcode.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (40.18%)
 * Total Accepted:    224.3K
 * Total Submissions: 558.2K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * Given a binary tree and a sum, find all root-to-leaf paths where each path's
 * sum equals the given sum.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given the below binary tree and sum = 22,
 *
 *
 * ⁠     5
 * ⁠    / \
 * ⁠   4   8
 * ⁠  /   / \
 * ⁠ 11  13  4
 * ⁠/  \    / \
 * 7    2  5   1
 *
 *
 * Return:
 *
 *
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
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
func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	sumIter(root, []int{},0, sum, &result)
	return result
}

func sumIter(root *TreeNode, alreadyIn []int, sum, target int, result *[][]int) {
	if root != nil {
		alreadyIn = append(alreadyIn, root.Val)
    sum = sum + root.Val
		if root.Left == nil && root.Right == nil {
       if sum == target{ 
				*result = append(*result, append([]int{}, alreadyIn...))
	    	} 
    }else {
			sumIter(root.Left, alreadyIn, sum, target, result)
			sumIter(root.Right, alreadyIn, sum, target, result)
		}
	}
}
