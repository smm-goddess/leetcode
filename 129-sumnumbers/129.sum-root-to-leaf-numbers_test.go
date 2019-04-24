package sumnumbers

import "testing"

func TestSumNumbers(t *testing.T) {
	tree := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	if sumNumbers(tree) != 25 {
		t.Error("First Error")
	}
	tree = &TreeNode{4, &TreeNode{9, &TreeNode{5, nil, nil}, &TreeNode{1, nil, nil}}, &TreeNode{0, nil, nil}}
	if sumNumbers(tree) != 1026 {
		t.Error("Second Error")
	}
}
