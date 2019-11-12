package find_bottom_left_tree_value

import "testing"

func TestFindBottomLeftTreeValue(t *testing.T) {
	tree := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	if findBottomLeftValue(tree) != 1 {
		t.Error("error")
	}
	tree = &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, nil, nil},
			nil},
		&TreeNode{3,
			&TreeNode{5,
				&TreeNode{7, nil, nil},
				nil},
			&TreeNode{6, nil, nil}}}
	if findBottomLeftValue(tree) != 7 {
		t.Error("7 error")
	}
}
