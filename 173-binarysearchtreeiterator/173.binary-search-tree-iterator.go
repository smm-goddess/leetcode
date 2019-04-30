package binarysearchtreeiterator

/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
 *
 * https://leetcode.com/problems/binary-search-tree-iterator/description/
 *
 * algorithms
 * Medium (47.84%)
 * Total Accepted:    199.4K
 * Total Submissions: 413.6K
 * Testcase Example:  '["BSTIterator","next","next","hasNext","next","hasNext","next","hasNext","next","hasNext"]\n[[[7,3,15,null,null,9,20]],[null],[null],[null],[null],[null],[null],[null],[null],[null]]'
 *
 * Implement an iterator over a binary search tree (BST). Your iterator will be
 * initialized with the root node of a BST.
 * 
 * Calling next() will return the next smallest number in the BST.
 * 
 * 
 * 
 * 
 * 
 * 
 * Example:
 * 
 * 
 * 
 * 
 * BSTIterator iterator = new BSTIterator(root);
 * iterator.next();    // return 3
 * iterator.next();    // return 7
 * iterator.hasNext(); // return true
 * iterator.next();    // return 9
 * iterator.hasNext(); // return true
 * iterator.next();    // return 15
 * iterator.hasNext(); // return true
 * iterator.next();    // return 20
 * iterator.hasNext(); // return false
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * next() and hasNext() should run in average O(1) time and uses O(h) memory,
 * where h is the height of the tree.
 * You may assume that next() call will always be valid, that is, there will be
 * at least a next smallest number in the BST when next() is called.
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
type BSTIterator struct {
	flatedBST *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{flatBST(root)}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	node := this.flatedBST
	this.flatedBST = this.flatedBST.Right
	node.Right = nil
	return node.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.flatedBST != nil
}

func flatBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	leftFlat, rightFlat := flatBST(root.Left), flatBST(root.Right)
	root.Left = nil
	root.Right = rightFlat
	if leftFlat == nil {
		return root
	} else {
		rightFlat = leftFlat
		for rightFlat.Right != nil {
			rightFlat = rightFlat.Right
		}
		rightFlat.Right = root
		return leftFlat
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
