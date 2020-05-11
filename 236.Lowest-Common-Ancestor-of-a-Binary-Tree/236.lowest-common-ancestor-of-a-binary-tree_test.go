package _36_Lowest_Common_Ancestor_of_a_Binary_Tree

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	root := &TreeNode{3,
		&TreeNode{5,
			&TreeNode{6, nil, nil},
			&TreeNode{2,
				&TreeNode{7, nil, nil},
				&TreeNode{4, nil, nil}}},
		&TreeNode{1,
			&TreeNode{0, nil, nil},
			&TreeNode{8, nil, nil}}}
	//fmt.Println(locateAncestor(root, &TreeNode{0, nil, nil}))
	fmt.Println(lowestCommonAncestor(root, &TreeNode{6, nil, nil}, &TreeNode{4, nil, nil}))
}
