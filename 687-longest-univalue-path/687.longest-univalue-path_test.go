package longest_univalue_path

import (
	"fmt"
	"testing"
)

func TestLongestUnivaluePath(t *testing.T) {
	tree := &TreeNode{5,
		&TreeNode{4,
			&TreeNode{1, nil, nil},
			&TreeNode{1, nil, nil}},
		&TreeNode{5, nil, &TreeNode{5, nil, nil}}}
	if longestUnivaluePath(tree) != 2 {
		t.Error(fmt.Sprintf("error"))
	}
	tree = &TreeNode{1,
		&TreeNode{4,
			&TreeNode{4, nil, nil},
			&TreeNode{4, nil, nil}},
		&TreeNode{5, nil, &TreeNode{5, nil, nil}}}
	if longestUnivaluePath(tree) != 2 {
		t.Error(fmt.Sprintf("error 2"))
	}
}
