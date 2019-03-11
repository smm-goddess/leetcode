package buildtree

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	tree := &TreeNode{
		3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil}},
	}
	//if !equalTree(tree, buildTree(preorder, inorder)) {
	//	t.Error("Error")
	//}

	preorder = []int{3, 9, 8, 1, 20, 15, 7}
	inorder = []int{8, 9, 1, 3, 15, 20, 7}
	tree = &TreeNode{
		3,
		&TreeNode{
			9,
			&TreeNode{8, nil, nil},
			&TreeNode{1, nil, nil},
		},
		&TreeNode{
			20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil},
		},
	}
	if !equalTree(tree, buildTree(preorder, inorder)) {
		t.Error("Error")
	}
	preorder = []int{1, 2}
	inorder = []int{1, 2}
	tree = &TreeNode{
		1,
		nil,
		&TreeNode{2, nil, nil},
	}
	if !equalTree(tree, buildTree(preorder, inorder)) {
		t.Error("Error")
	}
}

func equalTree(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	} else {
		return a.Val == b.Val && equalTree(a.Right, b.Right) && equalTree(a.Left, b.Left)
	}
}
