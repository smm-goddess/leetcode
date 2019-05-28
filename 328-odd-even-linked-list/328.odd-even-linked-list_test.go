package oddevenlist

import (
	"testing"
)

func listEqual(a, b *ListNode) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else {
		return a.Val == b.Val && listEqual(a.Next, b.Next)
	}
}

func TestOddEvenList(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	answer := &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{2, &ListNode{4, nil}}}}}
	if !listEqual(oddEvenList(list), answer) {
		t.Error("error")
	}
}
