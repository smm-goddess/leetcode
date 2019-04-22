package pathsum

import (
	"testing"
)

func sliceEqual(a, b []int) bool {
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

func TestPathSum(t *testing.T) {
	tree := &TreeNode{5,
		&TreeNode{4,
			&TreeNode{11,
				&TreeNode{7, nil, nil},
				&TreeNode{2, nil, nil}},
			nil},
		&TreeNode{8,
			&TreeNode{13, nil, nil},
			&TreeNode{4,
				&TreeNode{5, nil, nil},
				&TreeNode{1, nil, nil}}}}
	paths := pathSum(tree, 22)
	if !sliceEqual(paths[0], []int{5, 4, 11, 2}) || !sliceEqual(paths[1], []int{5, 8, 4, 5}) {
		t.Error("Test Fail")
	}
	pathSum(&TreeNode{1, &TreeNode{2, nil, nil}, nil}, 1)
}
