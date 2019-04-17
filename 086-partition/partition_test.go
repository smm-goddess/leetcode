package partition

import "testing"

func listEquals(h1, h2 *ListNode) bool {
	head1 := h1
	head2 := h2
	for head1 != nil && head2 != nil {
		if head1.Val != head2.Val {
			return false
		}
		head1 = head1.Next
		head2 = head2.Next
	}
	return head1 == nil && head2 == nil
}

func TestPartition(t *testing.T) {
	head := &ListNode{1,
		&ListNode{4,
			&ListNode{3,
				&ListNode{2,
					&ListNode{5,
						&ListNode{2, nil}}}}}}
	head = partition(head, 3)
	if !listEquals(head, &ListNode{1,
		&ListNode{2,
			&ListNode{2,
				&ListNode{4,
					&ListNode{3,
						&ListNode{5, nil}}}}}}) {
		t.Error("Error")
	}
}
