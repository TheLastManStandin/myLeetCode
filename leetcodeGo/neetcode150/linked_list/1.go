package linked_list

import "sync"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	var next *ListNode
	var prev *ListNode
	stop := false
	once := sync.Once{}
	for !stop {
		if head.Next != nil {
			next = head.Next
		} else {
			stop = true
		}
		if prev != nil {
			head.Next = prev
		}
		prev = head
		head = next
		once.Do(func() {
			prev.Next = nil
		})
	}
	return head
}
