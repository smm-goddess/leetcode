package preordertraversal

import (
	"testing"
)

func listEquals(a, b []int) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil || len(a) != len(b) {
		return false
	} else {
		for i, v := range a {
			if b[i] != v {
				return false
			}
		}
		return true
	}
}

func TestPreorderTraversal(t *testing.T) {
	tree := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	traversal := preorderTraversal(tree)
	if !listEquals(traversal, []int{1, 2, 3}) {
		t.Error("Traversal Error ")
	}
}
