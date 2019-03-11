package inordertraversal

import "testing"

func TestInorderTraversal(t *testing.T) {
	tree := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	traversalResult := inorderTraversal(tree)
	supose := []int{1, 3, 2}
	for i := 0; i < 3; i++ {
		if traversalResult[i] != supose[i] {
			t.Error("Error", traversalResult)
			break
		}
	}
}
