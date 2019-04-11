package rotateright

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	return strconv.Itoa(node.Val)
}
