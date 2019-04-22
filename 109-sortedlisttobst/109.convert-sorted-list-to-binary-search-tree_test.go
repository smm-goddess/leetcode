package sortedlisttobst

import (
	"fmt"
	"testing"
)

func treeEquals(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else {
		fmt.Printf("a:%d,b:%d\n", a.Val, b.Val)
		return a.Val == b.Val && treeEquals(a.Right, b.Right) && treeEquals(a.Left, b.Left)
	}
}

func preorder(a *TreeNode) {
	if a != nil {
		fmt.Printf(" %d", a.Val)
		preorder(a.Left)
		preorder(a.Right)
	}
}

func inorder(a *TreeNode) {
	if a != nil {
		inorder(a.Left)
		fmt.Printf(" %d", a.Val)
		inorder(a.Right)
	}
}

func TestSortedListToBST(t *testing.T) {
	list := &ListNode{-10, &ListNode{-3, &ListNode{0, &ListNode{5, &ListNode{9, nil}}}}}
	tree := sortedListToBST(list)
	tr := &TreeNode{0,
		&TreeNode{-3, &TreeNode{-10, nil, nil}, nil},
		&TreeNode{9, &TreeNode{5, nil, nil}, nil}}
	if !treeEquals(tree, tr) {
		t.Error("Test Fail")
	}
}
