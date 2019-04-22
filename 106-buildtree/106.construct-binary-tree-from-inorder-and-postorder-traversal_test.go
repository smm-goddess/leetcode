package buildtree

import (
	"testing"
)

func treeEquals(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else {
		return a.Val == b.Val && treeEquals(a.Left, b.Left) && treeEquals(a.Right, b.Right)
	}
}

func TestBuildTree(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	tree := &TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil}}}
	if !treeEquals(tree, buildTree(inorder, postorder)) {
		t.Error("Test Fail")
	}

}
