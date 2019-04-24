package flatten

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) String() string {
	return fmt.Sprintf("%d -> %s", node.Val, node.Right)
}
