package reversebetween

import "testing"

func TestReverseBetween(t *testing.T) {
	head := &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4,
					&ListNode{5, nil}}}}}
	reverseBetween(head, 5, 5)
	reverseBetween(&ListNode{1, &ListNode{2, nil}}, 1, 2)
}
