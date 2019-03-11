package validbst

import "testing"

func TestValidBST(t *testing.T) {
	tree := &TreeNode{
		2,
		&TreeNode{
			1, nil, nil,
		},
		&TreeNode{
			3, nil, nil,
		},
	}
	if !isValidBST(tree) {
		t.Error("1 2 3 fail")
	}
	tree = &TreeNode{
		5,
		&TreeNode{
			1,
			&TreeNode{
				3, nil, nil,
			},
			&TreeNode{6, nil, nil},
		},
		&TreeNode{
			4, nil, nil,
		},
	}
	if isValidBST(tree) {
		t.Error("5 1 4 nil nil 3 6 fail")
	}
	tree = &TreeNode{
		10,
		&TreeNode{
			5,
			&TreeNode{
				6, nil, nil,
			},
			&TreeNode{20, nil, nil},
		},
		&TreeNode{
			15, nil, nil,
		},
	}
	if isValidBST(tree) {
		t.Error("10 5 15 nil nil 6 20 fail")
	}
}
