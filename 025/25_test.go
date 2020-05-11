package reverse_group

import "testing"
import "fmt"

func TestReverseKGroup(t *testing.T) {
	list := &ListNode{1 , &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println(list)
	fmt.Println(reverseKGroup(list, 2))
}