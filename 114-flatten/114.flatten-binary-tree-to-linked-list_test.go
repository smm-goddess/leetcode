package flatten

import (
	"testing"
)

func TestFlatten(t *testing.T) {
	tree := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{3, nil, nil},
			&TreeNode{4, nil, nil}},
		&TreeNode{5,
			nil,
			&TreeNode{6, nil, nil}}}
	flatten(tree)
}
