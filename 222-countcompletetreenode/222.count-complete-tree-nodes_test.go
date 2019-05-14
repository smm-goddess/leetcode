package countnode

import "testing"

func TestCountNode(t *testing.T) {
	tree := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil}},
		&TreeNode{3,
			&TreeNode{6, nil, nil},
			nil}}
	if countNodes(tree) != 6 {
		t.Error("Fail")
	}
}
