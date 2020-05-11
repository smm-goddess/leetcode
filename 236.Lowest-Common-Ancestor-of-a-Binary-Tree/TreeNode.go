package _36_Lowest_Common_Ancestor_of_a_Binary_Tree

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) String() string {
	return strconv.Itoa(node.Val)
}
