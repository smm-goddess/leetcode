package issubtree

import (
	"testing"
)

func TestIsSubtree(t *testing.T) {
	t.Log("start issubtree test")
	t1 := &TreeNode{
		3,
		&TreeNode{
			4,
			&TreeNode{1, nil, nil},
			&TreeNode{2, nil, nil},
		},
		&TreeNode{5, nil, nil},
	}
	t2 := &TreeNode{
		4,
		&TreeNode{1, nil, nil},
		&TreeNode{2, nil, nil},
	}
	if !isSubtree(t1, t2) {
		t.Error("subtree test fail")
	}
	t1 = &TreeNode{
		3,
		&TreeNode{
			4,
			&TreeNode{1, nil, nil},
			&TreeNode{2,
				&TreeNode{0, nil, nil},
				nil,
			},
		},
		&TreeNode{5, nil, nil},
	}
	if isSubtree(t1, t2) {
		t.Error("subtree test fail")
	}
	t.Log("end issubtree test")
}
