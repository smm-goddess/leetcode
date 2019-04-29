package reorderlist

import (
	"testing"
)

func listEquals(l, b *ListNode) bool {
	for l != nil && b != nil {
		if l.Val != b.Val {
			return false
		}
		l, b = l.Next, b.Next
	}
	return l == nil && b == nil
}

func TestReorderList(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	list := &ListNode{1, &ListNode{4, &ListNode{2, &ListNode{3, nil}}}}
	reorderList(head)
	if !listEquals(head, list) {
		t.Error("First Fail")
	}
	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	list = &ListNode{1, &ListNode{5, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}}
	reorderList(head)
	if !listEquals(head, list) {
		t.Error("Second Fail")
	}
}
