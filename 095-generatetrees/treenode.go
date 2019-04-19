package generatetrees

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) String() string {
	return fmt.Sprintf(
		`
     %d
    /  \
  %s    %s
`, node.Val, node.Left, node.Right)
}
