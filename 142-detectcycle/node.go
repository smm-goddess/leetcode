package detectcycle

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	return fmt.Sprintf("%d", n.Val)
}
