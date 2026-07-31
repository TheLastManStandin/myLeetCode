package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	dummy := result

	if list1 == nil && list2 != nil {
		dummy = list2
		return dummy
	} else if list1 != nil && list2 == nil {
		dummy = list1
		return dummy
	} else if list1 == nil && list2 == nil {
		return nil
	}

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			dummy.Next = list1
			dummy = dummy.Next
			list1 = list1.Next
		} else {
			dummy.Next = list2
			dummy = dummy.Next
			list2 = list2.Next
		}
	}

	if list1 == nil && list2 != nil {
		dummy.Next = list2
	} else if list1 != nil && list2 == nil {
		dummy.Next = list1
	}

	return result.Next
}
