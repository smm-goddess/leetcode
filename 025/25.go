package reverse_group

import "fmt"

func reverse(head *ListNode) *ListNode{
	var h *ListNode
	for head != nil {
		node := head
		if h == nil {
			h = node
			head = head.Next
			h.Next = nil
		} else {
			head = head.Next
			node.Next = h
			h = node
		}
	}
	return h
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	cnt := 0
	var h, t,  tempList,tempListTail *ListNode
	for head != nil {
		cnt++
		node := head
		head = head.Next
		node.Next = nil
		if tempList == nil {
			tempList = node
			tempListTail = node
		} else {
			tempListTail.Next = node
			tempListTail = node
		}
		fmt.Println(tempList)
		if cnt == k  {
			reversed := reverse(tempList)
			if h == nil {
				h = reversed
				t = reversed
			} else {
				t.Next = reversed
			}
			for t.Next != nil {
				t = t.Next
			}
			tempListTail = nil
			tempList = nil
			cnt = 0
		}
	}

	if tempList != nil {
		if h == nil {
			h = tempList
		} else {
			t.Next = tempList
		}
	}
	
	return h
}