package deleteduplicates

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	return fmt.Sprintf("%d --> %s", node.Val, node.Next)
}
