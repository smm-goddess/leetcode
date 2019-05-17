package kthsmallest

import "testing"

func TestKthSmallest(t *testing.T) {
	root := &TreeNode{3,
		&TreeNode{1,
			nil,
			&TreeNode{2, nil, nil}},
		&TreeNode{4, nil, nil}}
	if kthSmallest(root, 1) != 1 {
		t.Error("! Error")
	}
	root = &TreeNode{5,
		&TreeNode{3,
			&TreeNode{2,
				&TreeNode{1, nil, nil},
				nil},
			&TreeNode{4, nil, nil}},
		&TreeNode{6, nil, nil}}
	if kthSmallest(root, 3) != 3 {
		t.Error("3 Error")
	}
}
