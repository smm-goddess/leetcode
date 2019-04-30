package sortlist

import "testing"

func listEquals(a, b *ListNode) bool {
	if a == nil || b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else {
		return a.Val == b.Val && listEquals(a.Next, b.Next)
	}
}
func TestSortList(t *testing.T) {
	list := &ListNode{4,
		&ListNode{2,
			&ListNode{1,
				&ListNode{3, nil}}}}
	list = sortList(list)
	if !listEquals(list,
		&ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4, nil}}}}) {
		t.Error("First Fail")
	}
	list = &ListNode{-1,
		&ListNode{5,
			&ListNode{3,
				&ListNode{4,
					&ListNode{0,
						&ListNode{4, nil}}}}}}
	list = sortList(list)
	if !listEquals(list,
		&ListNode{-1,
			&ListNode{0,
				&ListNode{3,
					&ListNode{4,
						&ListNode{4,
							&ListNode{5, nil}}}}}}) {
		t.Error("Second Fail")
	}
}
