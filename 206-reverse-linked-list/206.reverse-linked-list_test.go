package reverse_linked_list

import "testing"

func listNodeEquals(la, lb *ListNode) bool {
	if la == nil && lb == nil {
		return true
	}
	ha, hb := la, lb
	for ha != nil && hb != nil {
		if ha.Val != hb.Val {
			return false
		}
		ha = ha.Next
		hb = hb.Next
	}
	return ha == nil && hb == nil
}

func TestListEquals(t *testing.T) {
	la := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	lb := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	if !listNodeEquals(la, lb) {
		t.Error("list equal func error")
	}
	lc := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	if listNodeEquals(la, lc) {
		t.Error("list equals func error 2")
	}
}

func Test_reverseList(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	answer := &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}
	if !listNodeEquals(reverseList(list), answer) {
		t.Error("reverse func error")
	}
}
