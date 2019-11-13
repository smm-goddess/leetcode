package minimum_absolute_difference_in_bst

import "testing"

func TestMinimumDifference(t *testing.T) {
	tree := &TreeNode{1, nil,
		&TreeNode{3, &TreeNode{2, nil, nil}, nil}}
	if getMinimumDifference(tree) != 1 {
		t.Error("error")
	}
	//[236,104,701,null,227,null,911]
	tree = &TreeNode{236,
		&TreeNode{104, nil, &TreeNode{227, nil, nil}},
		&TreeNode{701, nil, &TreeNode{911, nil, nil}}}
	if getMinimumDifference(tree) != 9 {
		t.Error("9 error")
	}
}
