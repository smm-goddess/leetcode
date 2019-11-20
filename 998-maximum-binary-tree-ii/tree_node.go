package insert_into_max_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree_equals(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	} else {
		return a.Val == b.Val && tree_equals(a.Right, b.Right) && tree_equals(a.Left, b.Left)
	}
}
