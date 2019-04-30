package binarysearchtreeiterator

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) String() string {
	return fmt.Sprintf("%d", node.Val)
}
